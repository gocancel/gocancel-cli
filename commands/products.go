package commands

import (
	"github.com/gocancel/gocancel-cli/commands/displayers"
	"github.com/spf13/cobra"
)

func newProductsCmd() *Command {
	cmd := &Command{
		Command: &cobra.Command{
			Use:     "products",
			Aliases: []string{"product", "p"},
			Short:   "Display commands for working with products",
			Long:    "The subcommands of `gocancel products` view any GoCancel products.",
		},
	}

	CmdBuilder(
		cmd,
		runProductsGet,
		"get <product-id>",
		"Get a product",
		`Get a product with the provided id.

Only basic information is included with the text output format. For complete product details, use the JSON format.`,
		writer,
		aliasOpt("g"),
		displayerType(&displayers.Products{}),
	)

	return cmd
}

// runProductsGet gets a product.
func runProductsGet(c *CmdConfig) error {
	if len(c.Args) < 1 {
		return NewMissingArgsErr(c.NS)
	}

	productID := c.Args[0]

	product, err := c.Products().Get(productID)
	if err != nil {
		return err
	}

	return c.Display(displayers.Products{product})
}
