package commands

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/gocancel/gocancel-go"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLettersCommand(t *testing.T) {
	cmd := newLettersCmd()
	assert.NotNil(t, cmd)
	assertCommandNames(t, cmd, "list", "create", "get")
}

func TestLettersList(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		letters := []*gocancel.Letter{{
			ID:                    gocancel.String("26468553-08bb-47c4-a28c-d80dec6ef3b2"),
			AccountID:             gocancel.String("f172758f-7718-41f4-95d6-d3fd931e0326"),
			OrganizationID:        gocancel.String("f172758f-7718-41f4-95d6-d3fd931e0326"),
			OrganizationName:      gocancel.String("Foo"),
			ProductID:             gocancel.String("f172758f-7718-41f4-95d6-d3fd931e0326"),
			ProductName:           gocancel.String("Bar"),
			ProviderID:            gocancel.String("f172758f-7718-41f4-95d6-d3fd931e0326"),
			ProviderConfiguration: &gocancel.ProviderConfiguration{"foo": "bar"},
			Locale:                gocancel.String("nl-NL"),
			State:                 gocancel.String("generating"),
			ProofOfIDs:            []*string{gocancel.String("1d7e5cf6-a871-48cd-b98a-1ecc6acbda96"), gocancel.String("0971e527-ea0d-4ba2-a87b-0e8e8d4f83a2")},
			Email:                 gocancel.String("cancellations@foo.com"),
			Fax:                   gocancel.String("71-336-4530"),
			Parameters:            &gocancel.LetterParameters{"foo": "bar"},
			SignatureType:         gocancel.String("text"),
			SignatureData:         gocancel.String("John Doe"),
			Metadata:              &gocancel.AccountMetadata{"foo": "bar"},
			CreatedAt:             &gocancel.Timestamp{Time: time.Date(2021, time.May, 27, 11, 49, 05, 0, time.UTC)},
			UpdatedAt:             &gocancel.Timestamp{Time: time.Date(2021, time.May, 27, 11, 49, 05, 0, time.UTC)},
		}}

		tm.letters.EXPECT().List().Times(1).Return(letters, nil)

		err := runLettersList(config)
		assert.NoError(t, err)
	})
}

func TestLettersCreate(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		request := &gocancel.LetterRequest{
			OrganizationID: "f172758f-7718-41f4-95d6-d3fd931e0326",
			Locale:         "nl-NL",
		}

		payloadFile, err := ioutil.TempFile("", "payload")
		require.NoError(t, err)
		defer func() {
			os.Remove(payloadFile.Name())
			payloadFile.Close()
		}()

		err = json.NewEncoder(payloadFile).Encode(&request)
		require.NoError(t, err)

		letter := &gocancel.Letter{
			ID:                    gocancel.String("26468553-08bb-47c4-a28c-d80dec6ef3b2"),
			AccountID:             gocancel.String("f172758f-7718-41f4-95d6-d3fd931e0326"),
			OrganizationID:        gocancel.String("f172758f-7718-41f4-95d6-d3fd931e0326"),
			OrganizationName:      gocancel.String("Foo"),
			ProductID:             gocancel.String("f172758f-7718-41f4-95d6-d3fd931e0326"),
			ProductName:           gocancel.String("Bar"),
			ProviderID:            gocancel.String("f172758f-7718-41f4-95d6-d3fd931e0326"),
			ProviderConfiguration: &gocancel.ProviderConfiguration{"foo": "bar"},
			Locale:                gocancel.String("nl-NL"),
			State:                 gocancel.String("generating"),
			ProofOfIDs:            []*string{gocancel.String("1d7e5cf6-a871-48cd-b98a-1ecc6acbda96"), gocancel.String("0971e527-ea0d-4ba2-a87b-0e8e8d4f83a2")},
			Email:                 gocancel.String("cancellations@foo.com"),
			Fax:                   gocancel.String("71-336-4530"),
			Parameters:            &gocancel.LetterParameters{"foo": "bar"},
			SignatureType:         gocancel.String("text"),
			SignatureData:         gocancel.String("John Doe"),
			Metadata:              &gocancel.AccountMetadata{"foo": "bar"},
			CreatedAt:             &gocancel.Timestamp{Time: time.Date(2021, time.May, 27, 11, 49, 05, 0, time.UTC)},
			UpdatedAt:             &gocancel.Timestamp{Time: time.Date(2021, time.May, 27, 11, 49, 05, 0, time.UTC)},
		}

		tm.letters.EXPECT().Create(request).Times(1).Return(letter, nil)

		viper.Set(nskey(config.NS, "payload"), payloadFile.Name())

		err = runLettersCreate(config)
		assert.NoError(t, err)
	})
}

func TestLettersGet(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		letter := &gocancel.Letter{
			ID:                    gocancel.String("26468553-08bb-47c4-a28c-d80dec6ef3b2"),
			AccountID:             gocancel.String("f172758f-7718-41f4-95d6-d3fd931e0326"),
			OrganizationID:        gocancel.String("f172758f-7718-41f4-95d6-d3fd931e0326"),
			OrganizationName:      gocancel.String("Foo"),
			ProductID:             gocancel.String("f172758f-7718-41f4-95d6-d3fd931e0326"),
			ProductName:           gocancel.String("Bar"),
			ProviderID:            gocancel.String("f172758f-7718-41f4-95d6-d3fd931e0326"),
			ProviderConfiguration: &gocancel.ProviderConfiguration{"foo": "bar"},
			Locale:                gocancel.String("nl-NL"),
			State:                 gocancel.String("generating"),
			ProofOfIDs:            []*string{gocancel.String("1d7e5cf6-a871-48cd-b98a-1ecc6acbda96"), gocancel.String("0971e527-ea0d-4ba2-a87b-0e8e8d4f83a2")},
			Email:                 gocancel.String("cancellations@foo.com"),
			Fax:                   gocancel.String("71-336-4530"),
			Parameters:            &gocancel.LetterParameters{"foo": "bar"},
			SignatureType:         gocancel.String("text"),
			SignatureData:         gocancel.String("John Doe"),
			Metadata:              &gocancel.AccountMetadata{"foo": "bar"},
			CreatedAt:             &gocancel.Timestamp{Time: time.Date(2021, time.May, 27, 11, 49, 05, 0, time.UTC)},
			UpdatedAt:             &gocancel.Timestamp{Time: time.Date(2021, time.May, 27, 11, 49, 05, 0, time.UTC)},
		}

		tm.letters.EXPECT().Get(*letter.ID).Times(1).Return(letter, nil)

		config.Args = append(config.Args, *letter.ID)

		err := runLettersGet(config)
		assert.NoError(t, err)
	})
}
