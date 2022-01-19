package displayers

import (
	"fmt"
	"io"
	"reflect"
	"strings"
	"text/tabwriter"
)

// Displayable is a displayable entity. These are used for printing results.
type Displayable interface {
	Cols() []string
	ColMap() map[string]string
	KV() []map[string]interface{}
	JSON(io.Writer) error
}

// Displayer has the display options, the item to display, and where to display to.
type Displayer struct {
	OutputType string
	ColumnList string
	NoHeaders  bool

	Item Displayable
	Out  io.Writer
}

// Display ends up rendering the content in one of two formats (text|json)
func (d *Displayer) Display() error {
	switch d.OutputType {
	case "json":
		if containsOnlyNilSlice(d.Item) {
			_, err := d.Out.Write([]byte("[]"))
			return err
		}
		return d.Item.JSON(d.Out)
	case "text":
		var cols []string
		for _, c := range strings.Split(strings.Join(strings.Fields(d.ColumnList), ""), ",") {
			if c != "" {
				cols = append(cols, c)
			}
		}

		return DisplayText(d.Item, d.Out, d.NoHeaders, cols)
	default:
		return fmt.Errorf("unknown output type")
	}
}

// DisplayText writes tabbed content to the passed in io.Writer
// while potentially adding or removing headers.
func DisplayText(item Displayable, out io.Writer, noHeaders bool, includeCols []string) error {
	w := new(tabwriter.Writer)
	w.Init(out, 0, 0, 4, ' ', 0)

	cols := item.Cols()
	if len(includeCols) > 0 && includeCols[0] != "" {
		cols = includeCols
	}

	if !noHeaders {
		headers := make([]string, 0, len(cols))
		for _, k := range cols {
			col := item.ColMap()[k]
			if col == "" {
				return fmt.Errorf("unknown column %q", k)
			}

			headers = append(headers, col)
		}
		fmt.Fprintln(w, strings.Join(headers, "\t"))
	}

	for _, r := range item.KV() {
		values := make([]interface{}, 0, len(cols))
		formats := make([]string, 0, len(cols))

		for _, col := range cols {
			v := r[col]

			values = append(values, v)

			switch v.(type) {
			case string:
				formats = append(formats, "%s")
			case int:
				formats = append(formats, "%d")
			case float64:
				formats = append(formats, "%f")
			case bool:
				formats = append(formats, "%v")
			default:
				formats = append(formats, "%v")
			}
		}
		format := strings.Join(formats, "\t")
		fmt.Fprintf(w, format+"\n", values...)
	}

	return w.Flush()
}

// containsOnlyNiSlice returns true if the given interface's concrete type is
// a pointer to a struct that contains a single nil slice field.
func containsOnlyNilSlice(i interface{}) bool {
	if reflect.TypeOf(i).Kind() != reflect.Ptr {
		return false
	}

	element := reflect.ValueOf(i).Elem()
	if element.NumField() != 1 {
		return false
	}

	slice := element.Field(0)
	if slice.Kind() != reflect.Slice {
		return false
	}

	if slice.Cap() != 0 {
		return false
	}
	if slice.Len() != 0 {
		return false
	}
	if slice.Pointer() != 0 {
		return false
	}

	return true
}
