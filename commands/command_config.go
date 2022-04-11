package commands

import (
	"fmt"
	"io"

	"github.com/gocancel/gocancel-cli/client"
	"github.com/gocancel/gocancel-cli/commands/displayers"
	"github.com/spf13/viper"
)

// CmdConfig is a command configuration.
type CmdConfig struct {
	NS   string
	Out  io.Writer
	Args []string

	initServices func(*CmdConfig) error

	// Services
	Categories    func() client.CategoriesService
	Letters       func() client.LettersService
	Organizations func() client.OrganizationsService
	Products      func() client.ProductsService
}

// NewCmdConfig creates an instance of a CmdConfig.
func NewCmdConfig(ns string, out io.Writer, args []string, initServices bool) (*CmdConfig, error) {
	cmdConfig := &CmdConfig{
		NS:   ns,
		Out:  out,
		Args: args,

		initServices: func(c *CmdConfig) error {
			clientID := viper.GetString("client-id")
			clientSecret := viper.GetString("client-secret")

			gocancelClient, err := client.Create(clientID, clientSecret)
			if err != nil {
				return fmt.Errorf("unable to initialize GoCancel API client: %s", err)
			}

			c.Categories = func() client.CategoriesService { return client.NewCategoriesService(gocancelClient) }
			c.Letters = func() client.LettersService { return client.NewLettersService(gocancelClient) }
			c.Organizations = func() client.OrganizationsService { return client.NewOrganizationsService(gocancelClient) }
			c.Products = func() client.ProductsService { return client.NewProductsService(gocancelClient) }

			return nil
		},
	}

	if initServices {
		if err := cmdConfig.initServices(cmdConfig); err != nil {
			return nil, err
		}
	}

	return cmdConfig, nil
}

// Display displays the output from a command.
func (c *CmdConfig) Display(d displayers.Displayable) error {
	dc := &displayers.Displayer{
		Item: d,
		Out:  c.Out,
	}

	columnList := viper.GetString(nskey(c.NS, "format"))
	withHeaders := viper.GetBool(nskey(c.NS, "no-header"))

	dc.NoHeaders = withHeaders
	dc.ColumnList = columnList
	dc.OutputType = output

	return dc.Display()
}
