package commands

import (
	"github.com/gocancel/gocancel-cli/commands/displayers"
	"github.com/gocancel/gocancel-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

	list := CmdBuilder(
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
	AddStringFlag(list, "slug", "", "", "A slug to filter organizations on.")
	AddStringFlag(list, "category", "", "", "A category to filter organizations on.")
	AddStringFlag(list, "url", "", "", "An (partial) URL to filter organizations on.")
	AddStringSliceFlag(list, "locales", "", []string{}, "One or more locales to filter organizations on.")

	CmdBuilder(
		cmd,
		runOrganizationsGet,
		"get <organization-id>",
		"Get an organization",
		`Get an organization with the provided id.

Only basic information is included with the text output format. For complete organization details, use the JSON format.`,
		writer,
		aliasOpt("g"),
		displayerType(&displayers.Categories{}),
	)

	listProducts := CmdBuilder(
		cmd,
		runOrganizationsListProducts,
		"list-products <organization-id>",
		"List all products for an organization",
		`List all products for an organization within your account.

Only basic information is included with the text output format. For complete product details, use the JSON format.`,
		writer,
		aliasOpt("lsp"),
		displayerType(&displayers.Products{}),
	)
	AddStringFlag(listProducts, "slug", "", "", "A slug to filter products on.")
	AddStringFlag(listProducts, "url", "", "", "An (partial) URL to filter products on.")
	AddStringSliceFlag(listProducts, "locales", "", []string{}, "One or more locales to filter products on.")

	CmdBuilder(
		cmd,
		runOrganizationsGetProduct,
		"get-product <organization-id> <product-id>",
		"Get a single product for an organization",
		`Get a single product for an organization within your account.

Only basic information is included with the text output format. For complete product details, use the JSON format.`,
		writer,
		aliasOpt("gp"),
		displayerType(&displayers.Products{}),
	)

	return cmd
}

// runOrganizationsList lists all organizations.
func runOrganizationsList(c *CmdConfig) error {
	opts := &gocancel.OrganizationsListOptions{}

	slug := viper.GetString(nskey(c.NS, "slug"))
	if slug != "" {
		opts.Slug = slug
	}

	category := viper.GetString(nskey(c.NS, "category"))
	if category != "" {
		opts.Category = category
	}

	url := viper.GetString(nskey(c.NS, "url"))
	if url != "" {
		opts.URL = url
	}

	locales := viper.GetStringSlice(nskey(c.NS, "locales"))
	if len(locales) > 0 {
		opts.Locales = locales
	}

	organizations, err := c.Organizations().List(opts)
	if err != nil {
		return err
	}

	return c.Display(displayers.Organizations(organizations))
}

// runOrganizationsGet gets an organization.
func runOrganizationsGet(c *CmdConfig) error {
	if len(c.Args) < 1 {
		return NewMissingArgsErr(c.NS)
	}

	organizationID := c.Args[0]

	organization, err := c.Organizations().Get(organizationID)
	if err != nil {
		return err
	}

	return c.Display(displayers.Organizations{organization})
}

// runOrganizationsListProducts lists all products for an organization.
func runOrganizationsListProducts(c *CmdConfig) error {
	if len(c.Args) < 1 {
		return NewMissingArgsErr(c.NS)
	}

	organizationID := c.Args[0]
	opts := &gocancel.OrganizationProductsListOptions{}

	slug := viper.GetString(nskey(c.NS, "slug"))
	if slug != "" {
		opts.Slug = slug
	}

	url := viper.GetString(nskey(c.NS, "url"))
	if url != "" {
		opts.URL = url
	}

	locales := viper.GetStringSlice(nskey(c.NS, "locales"))
	if len(locales) > 0 {
		opts.Locales = locales
	}

	products, err := c.Organizations().ListProducts(organizationID, opts)
	if err != nil {
		return err
	}

	return c.Display(displayers.Products(products))
}

// runOrganizationsGetProduct lists all products for an organization.
func runOrganizationsGetProduct(c *CmdConfig) error {
	if len(c.Args) < 2 {
		return NewMissingArgsErr(c.NS)
	}

	organizationID := c.Args[0]
	productID := c.Args[1]

	product, err := c.Organizations().GetProduct(organizationID, productID)
	if err != nil {
		return err
	}

	return c.Display(displayers.Products{product})
}
