package commands

import (
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

	CmdBuilder(
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
	organizations, err := c.Organizations().List()
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

	products, err := c.Organizations().ListProducts(organizationID)
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
