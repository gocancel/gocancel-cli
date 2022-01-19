package commands

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/shiena/ansicolor"
	"github.com/spf13/viper"
)

var (
	colorErr    = color.RedString("Error")
	colorWarn   = color.YellowString("Warning")
	colorNotice = color.GreenString("Notice")

	// errAction specifies what should happen when an error occurs
	errAction = func() {
		os.Exit(1)
	}
)

func init() {
	color.Output = ansicolor.NewAnsiColorWriter(os.Stderr)
}

type outputErrors struct {
	Errors []outputError `json:"errors"`
}

type outputError struct {
	Detail string `json:"detail"`
}

func checkErr(err error) {
	if err == nil {
		return
	}

	output := viper.GetString("output")

	switch output {
	default:
		fmt.Fprintf(color.Output, "%s: %v\n", colorErr, err)
	case "json":
		es := outputErrors{
			Errors: []outputError{
				{Detail: err.Error()},
			},
		}

		b, _ := json.Marshal(&es)
		fmt.Println(string(b))
	}

	errAction()
}

func warn(msg string, args ...interface{}) {
	fmt.Fprintf(color.Output, "%s: %s\n", colorWarn, fmt.Sprintf(msg, args...))
}
func warnConfirm(msg string, args ...interface{}) {
	fmt.Fprintf(color.Output, "%s: %s", colorWarn, fmt.Sprintf(msg, args...))
}

func notice(msg string, args ...interface{}) {
	fmt.Fprintf(color.Output, "%s: %s\n", colorNotice, fmt.Sprintf(msg, args...))
}

// MissingArgsErr is returned when there are too few arguments for a command.
type MissingArgsErr struct {
	Command string
}

var _ error = &MissingArgsErr{}

// NewMissingArgsErr creates a MissingArgsErr instance.
func NewMissingArgsErr(cmd string) *MissingArgsErr {
	return &MissingArgsErr{Command: cmd}
}

func (e *MissingArgsErr) Error() string {
	return fmt.Sprintf("(%s) command is missing required arguments", e.Command)
}
