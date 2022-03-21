package displayers

import (
	"encoding/json"
	"io"

	"github.com/gocancel/gocancel-go"
)

type Categories []*gocancel.Category

var _ Displayable = (*Categories)(nil)

func (c Categories) Cols() []string {
	return []string{
		"ID",
		"Name",
		"Slug",
		"RequiresConsent",
		"Metadata",
		"CreatedAt",
		"UpdatedAt",
	}
}

func (c Categories) ColMap() map[string]string {
	return map[string]string{
		"ID":              "ID",
		"Name":            "Name",
		"Slug":            "Slug",
		"RequiresConsent": "Requires Consent",
		"Metadata":        "Metadata",
		"CreatedAt":       "Created At",
		"UpdatedAt":       "Updated At",
	}
}

func (c Categories) KV() []map[string]interface{} {
	out := make([]map[string]interface{}, len(c))

	for i, category := range c {
		out[i] = map[string]interface{}{
			"ID":              *category.ID,
			"Name":            *category.Name,
			"Slug":            *category.Slug,
			"RequiresConsent": *category.RequiresConsent,
			"Metadata":        *category.Metadata,
			"CreatedAt":       category.CreatedAt,
			"UpdatedAt":       category.UpdatedAt,
		}
	}

	return out
}

func (c Categories) JSON(w io.Writer) error {
	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	return e.Encode(c)
}
