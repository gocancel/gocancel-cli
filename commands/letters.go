package commands

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/gocancel/gocancel-cli/commands/displayers"
	"github.com/gocancel/gocancel-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vincent-petithory/dataurl"
	"sigs.k8s.io/yaml"
)

func newLettersCmd() *Command {
	cmd := &Command{
		Command: &cobra.Command{
			Use:     "letters",
			Aliases: []string{"letter", "l"},
			Short:   "Display commands for working with letters",
			Long:    "The subcommands of `gocancel letters` manage your GoCancel letters.",
		},
	}

	create := CmdBuilder(
		cmd,
		runLettersCreate,
		"create",
		"Create a letter",
		`Create a letter with the given payload.`,
		writer,
		aliasOpt("c"),
		displayerType(&displayers.Letters{}),
	)
	AddStringFlag(create, "payload", "", "", `Path to letter payload in JSON or YAML format. Set to "-" to read from stdin.`, requiredOpt())
	AddStringSliceFlag(create, "proof-of-id-file", "", []string{}, `Path to a proof of ID file to attach.`)

	CmdBuilder(
		cmd,
		runLettersGet,
		"get <letter id>",
		"Get a letter",
		`Get a letter with the provider id.

Only basic information is included with the text output format. For complete letter details, use the JSON format.`,
		writer,
		aliasOpt("g"),
		displayerType(&displayers.Letters{}),
	)

	CmdBuilder(
		cmd,
		runLettersList,
		"list",
		"List all letters",
		`List all letters within your account.

Only basic information is included with the text output format. For complete letter details, use the JSON format.`,
		writer,
		aliasOpt("ls"),
		displayerType(&displayers.Letters{}),
	)

	return cmd
}

// runLettersCreate creates an app.
func runLettersCreate(c *CmdConfig) error {
	payloadPath := viper.GetString(nskey(c.NS, "payload"))

	letterPayload, err := readLetterPayload(os.Stdin, payloadPath)
	if err != nil {
		return err
	}

	proofOfIDPaths := viper.GetStringSlice(nskey(c.NS, "proof-of-id-file"))

	proofOfIDs, err := encodeAttachments(proofOfIDPaths)
	if err != nil {
		return err
	}

	letterPayload.ProofOfIDs = proofOfIDs

	letter, _, err := c.Client.Letters.Create(context.Background(), letterPayload.CreateLetterRequest)
	if err != nil {
		return err
	}

	notice("Letter created")

	return c.Display(displayers.Letters{letter})
}

// runLettersGet gets a letter.
func runLettersGet(c *CmdConfig) error {
	if len(c.Args) < 1 {
		return NewMissingArgsErr(c.NS)
	}

	id := c.Args[0]

	letter, _, err := c.Client.Letters.Get(context.Background(), id)
	if err != nil {
		return err
	}

	return c.Display(displayers.Letters{letter})
}

// runLettersList lists all letters.
func runLettersList(c *CmdConfig) error {
	letters, _, err := c.Client.Letters.List(context.Background(), nil)
	if err != nil {
		return err
	}

	return c.Display(displayers.Letters(letters))
}

func readLetterPayload(stdin io.Reader, path string) (*LetterPayload, error) {
	var payload io.Reader
	if path == "-" {
		payload = stdin
	} else {
		payloadFile, err := os.Open(path)
		if err != nil {
			if os.IsNotExist(err) {
				return nil, fmt.Errorf("opening letter payload: %s does not exist", path)
			}
			return nil, fmt.Errorf("opening letter payload: %w", err)
		}

		defer payloadFile.Close()
		payload = payloadFile
	}

	b, err := ioutil.ReadAll(payload)
	if err != nil {
		return nil, fmt.Errorf("reading letter payload: %w", err)
	}

	p, err := parseLetterPayload(b)
	if err != nil {
		return nil, fmt.Errorf("parsing letter payload: %w", err)
	}

	return p, nil
}

type LetterPayload struct {
	*gocancel.CreateLetterRequest
}

func parseLetterPayload(payload []byte) (*LetterPayload, error) {
	jsonPayload, err := yaml.YAMLToJSON(payload)
	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(bytes.NewReader(jsonPayload))
	dec.DisallowUnknownFields()

	var letterPayload LetterPayload
	if err := dec.Decode(&letterPayload); err != nil {
		return nil, err
	}

	return &letterPayload, nil
}

func encodeAttachments(attachmentPaths []string) ([]string, error) {
	var attachments []string
	for _, attachmentPath := range attachmentPaths {
		attachment, err := encodeAttachment(attachmentPath)
		if err != nil {
			return nil, err
		}

		attachments = append(attachments, *attachment)
	}

	return attachments, nil
}

func encodeAttachment(attachmentPath string) (*string, error) {
	attachmentFile, err := ioutil.ReadFile(attachmentPath)
	if err != nil {
		return nil, err
	}

	attachment := dataurl.EncodeBytes(attachmentFile)

	return &attachment, nil
}
