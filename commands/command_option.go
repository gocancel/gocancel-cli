package commands

import (
	"github.com/gocancel/gocancel-cli/commands/displayers"
)

// cmdOption allow configuration of a command.
type cmdOption func(*Command)

// aliasOpt adds aliases for a command.
func aliasOpt(aliases ...string) cmdOption {
	return func(c *Command) {
		if c.Aliases == nil {
			c.Aliases = []string{}
		}

		c.Aliases = append(c.Aliases, aliases...)
	}
}

// displayerType sets the columns for display for a command.
func displayerType(d displayers.Displayable) cmdOption {
	return func(c *Command) {
		c.fmtCols = d.Cols()
	}
}

// hiddenCmd make a command hidden.
func hiddenCmd() cmdOption {
	return func(c *Command) {
		c.Hidden = true
	}
}
