package commands

import (
	"github.com/gocancel/gocancel-cli/commands/displayers"
	"github.com/gocancel/gocancel-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

	list := CmdBuilder(
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
	AddStringFlag(list, "slug", "", "", "A slug to filter categories on.")
	AddStringSliceFlag(list, "locales", "", []string{}, "One or more locales to filter categories on.")

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
	opts := &gocancel.CategoriesListOptions{}

	slug := viper.GetString(nskey(c.NS, "slug"))
	if slug != "" {
		opts.Slug = slug
	}

	locales := viper.GetStringSlice(nskey(c.NS, "locales"))
	if len(locales) > 0 {
		opts.Locales = locales
	}

	categories, err := c.Categories().List(opts)
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
