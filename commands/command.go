package commands

import (
	"fmt"
	"io"
	"strings"

	"github.com/spf13/cobra"
)

// Command is a wrapper around cobra.Command that adds gocancel specific
// functionality.
type Command struct {
	*cobra.Command

	fmtCols       []string
	childCommands []*Command
}

// AddCommand adds child commands and adds child commands for cobra as well.
func (c *Command) AddCommand(commands ...*Command) {
	c.childCommands = append(c.childCommands, commands...)
	for _, cmd := range commands {
		c.Command.AddCommand(cmd.Command)
	}
}

// CmdRunner runs a command and passes in a cmdConfig.
type CmdRunner func(*CmdConfig) error

func CmdBuilder(parent *Command, cr CmdRunner, cliText, shortdesc string, longdesc string, out io.Writer, options ...cmdOption) *Command {
	return cmdBuilderWithInit(parent, cr, cliText, shortdesc, longdesc, out, true, options...)
}

func cmdBuilderWithInit(parent *Command, cr CmdRunner, cliText, shortdesc string, longdesc string, out io.Writer, initCmd bool, options ...cmdOption) *Command {
	cc := &cobra.Command{
		Use:   cliText,
		Short: shortdesc,
		Long:  longdesc,
		Run: func(cmd *cobra.Command, args []string) {
			c, err := NewCmdConfig(
				cmdNS(cmd),
				out,
				args,
				initCmd,
			)
			checkErr(err)

			err = cr(c)
			checkErr(err)
		},
	}

	c := &Command{Command: cc}

	if parent != nil {
		parent.AddCommand(c)
	}

	for _, co := range options {
		co(c)
	}

	if cols := c.fmtCols; cols != nil {
		formatHelp := fmt.Sprintf("Columns for output in a comma-separated list. Possible values: `%s`", strings.Join(cols, "`"+", "+"`"))
		AddStringFlag(c, "format", "", "", formatHelp)
		AddBoolFlag(c, "no-header", "", false, "Return raw data with no headers")
	}

	return c
}
