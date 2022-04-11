package commands

import (
	"github.com/gocancel/gocancel-cli/commands/displayers"
	"github.com/spf13/cobra"
)

func newCategoriesCmd() *Command {
	cmd := &Command{
		Command: &cobra.Command{
			Use:     "categories",
			Aliases: []string{"category", "c"},
			Short:   "Display commands for working with categories",
			Long:    "The subcommands of `gocancel categories` view any GoCancel categories.",
		},
	}

	CmdBuilder(
		cmd,
		runCategoriesList,
		"list",
		"List all categories",
		`List all categories within your account.

Only basic information is included with the text output format. For complete category details, use the JSON format.`,
		writer,
		aliasOpt("ls"),
		displayerType(&displayers.Categories{}),
	)

	CmdBuilder(
		cmd,
		runCategoriesGet,
		"get <category-id>",
		"Get a category",
		`Get a category with the provided id.

Only basic information is included with the text output format. For complete category details, use the JSON format.`,
		writer,
		aliasOpt("g"),
		displayerType(&displayers.Categories{}),
	)

	return cmd
}

// runCategoriesList lists all categories.
func runCategoriesList(c *CmdConfig) error {
	categories, err := c.Categories().List()
	if err != nil {
		return err
	}

	return c.Display(displayers.Categories(categories))
}

// runCategoriesGet gets a category.
func runCategoriesGet(c *CmdConfig) error {
	if len(c.Args) < 1 {
		return NewMissingArgsErr(c.NS)
	}

	categoryID := c.Args[0]

	category, err := c.Categories().Get(categoryID)
	if err != nil {
		return err
	}

	return c.Display(displayers.Categories{category})
}
