package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// rootCmd is the root level gocancel command that all other commands attach to.
	rootCmd = &Command{ // base command
		Command: &cobra.Command{
			Use:   "gocancel",
			Short: "A CLI to interact with the GoCancel API.",
			Long:  "The official command-line tool to interact with the GoCancel API.",
		},
	}

	// output global output format.
	output string
	// writer wires up stdout for all commands to write to.
	writer = os.Stdout

	requiredColor = color.New(color.Bold).SprintfFunc()
)

// Execute executes the current command using rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	initConfig()

	rootPFlagSet := rootCmd.PersistentFlags()

	rootPFlagSet.StringVarP(&output, "output", "o", "text", "Desired output format [text|json]")
	_ = viper.BindPFlag("output", rootPFlagSet.Lookup("output"))

	rootCmd.AddCommand(newLettersCmd())
}

func initConfig() {
	viper.SetEnvPrefix("GOCANCEL")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	viper.SetDefault("output", "text")
}

type flagOpt func(c *Command, name, key string)

func requiredOpt() flagOpt {
	return func(c *Command, name, key string) {
		_ = c.MarkFlagRequired(key)

		key = fmt.Sprintf("required.%s", key)
		viper.Set(key, true)

		u := c.Flag(name).Usage
		c.Flag(name).Usage = fmt.Sprintf("%s %s", u, requiredColor("(required)"))
	}
}

// AddStringFlag adds a string flag to a command.
func AddStringFlag(cmd *Command, name, shorthand, dflt, desc string, opts ...flagOpt) {
	fn := flagName(cmd, name)
	cmd.Flags().StringP(name, shorthand, dflt, desc)

	for _, o := range opts {
		o(cmd, name, fn)
	}

	_ = viper.BindPFlag(fn, cmd.Flags().Lookup(name))
}

// AddBoolFlag adds a boolean flag to a command.
func AddBoolFlag(cmd *Command, name, shorthand string, def bool, desc string, opts ...flagOpt) {
	fn := flagName(cmd, name)
	cmd.Flags().BoolP(name, shorthand, def, desc)
	viper.BindPFlag(fn, cmd.Flags().Lookup(name))

	for _, o := range opts {
		o(cmd, name, fn)
	}
}

// AddStringSliceFlag adds a string slice flag to a command.
func AddStringSliceFlag(cmd *Command, name, shorthand string, def []string, desc string, opts ...flagOpt) {
	fn := flagName(cmd, name)
	cmd.Flags().StringSliceP(name, shorthand, def, desc)
	viper.BindPFlag(fn, cmd.Flags().Lookup(name))

	for _, o := range opts {
		o(cmd, name, fn)
	}
}

func flagName(cmd *Command, name string) string {
	if cmd.Parent() != nil {
		return fmt.Sprintf("%s.%s", cmd.Parent().Name(), nskey(cmd.Name(), name))
	}

	return nskey(cmd.Name(), name)
}

func cmdNS(cmd *cobra.Command) string {
	if cmd.Parent() != nil {
		return fmt.Sprintf("%s.%s", cmd.Parent().Name(), cmd.Name())
	}

	return cmd.Name()
}

func nskey(ns, key string) string {
	return fmt.Sprintf("%s.%s", ns, key)
}
