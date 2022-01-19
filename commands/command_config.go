package commands

import (
	"fmt"
	"io"

	"github.com/gocancel/gocancel-cli/client"
	"github.com/gocancel/gocancel-cli/commands/displayers"
	"github.com/gocancel/gocancel-go"
	"github.com/spf13/viper"
)

// CmdConfig is a command configuration.
type CmdConfig struct {
	NS     string
	Client *gocancel.Client
	Out    io.Writer
	Args   []string

	initClient func(*CmdConfig) error
}

// NewCmdConfig creates an instance of a CmdConfig.
func NewCmdConfig(ns string, out io.Writer, args []string, initClient bool) (*CmdConfig, error) {
	cmdConfig := &CmdConfig{
		NS:   ns,
		Out:  out,
		Args: args,

		initClient: func(c *CmdConfig) error {
			clientID := viper.GetString("client-id")
			clientSecret := viper.GetString("client-secret")

			client, err := client.Create(clientID, clientSecret)
			if err != nil {
				return fmt.Errorf("unable to initialize GoCancel API client: %s", err)
			}

			c.Client = client

			return nil
		},
	}

	if initClient {
		if err := cmdConfig.initClient(cmdConfig); err != nil {
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
