package commands

import (
	"context"

	"github.com/gocancel/gocancel-cli/commands/displayers"
	"github.com/spf13/cobra"
)

func newOrganizationsCmd() *Command {
	cmd := &Command{
		Command: &cobra.Command{
			Use:     "organizations",
			Aliases: []string{"organization", "o"},
			Short:   "Display commands for working with organizations",
			Long:    "The subcommands of `gocancel organizations` view any GoCancel organizations.",
		},
	}

	CmdBuilder(
		cmd,
		runOrganizationsList,
		"list",
		"List all organizations",
		`List all organizations within your account.

Only basic information is included with the text output format. For complete organization details, use the JSON format.`,
		writer,
		aliasOpt("ls"),
		displayerType(&displayers.Organizations{}),
	)

	return cmd
}

// runOrganizationsList lists all organizations.
func runOrganizationsList(c *CmdConfig) error {
	organizations, _, err := c.Client.Organizations.List(context.Background(), nil)
	if err != nil {
		return err
	}

	return c.Display(displayers.Organizations(organizations))
}
