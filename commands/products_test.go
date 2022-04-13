package commands

import (
	"testing"
	"time"

	"github.com/gocancel/gocancel-go"
	"github.com/stretchr/testify/assert"
)

func TestProductsCommand(t *testing.T) {
	cmd := newProductsCmd()
	assert.NotNil(t, cmd)
	assertCommandNames(t, cmd, "get")
}

func TestProductsGet(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		product := &gocancel.Product{
			ID:                gocancel.String("ad41acf6-9f3c-45cd-9438-134d6ee7831c"),
			Name:              gocancel.String("ACME"),
			Slug:              gocancel.String("acme"),
			Email:             gocancel.String("contact@acme.com"),
			URL:               gocancel.String("https://acme.com"),
			Phone:             gocancel.String("517-234-9141"),
			Fax:               nil,
			OrganizationID:    gocancel.String("f38c8fab-0fa6-40b6-bb0c-6b3dfa2fec05"),
			RequiresConsent:   gocancel.Bool(true),
			RequiresProofOfID: gocancel.Bool(true),
			Metadata:          &gocancel.AccountMetadata{"foo": "bar"},
			CreatedAt:         &gocancel.Timestamp{Time: time.Date(2021, time.May, 27, 11, 49, 05, 0, time.UTC)},
			UpdatedAt:         &gocancel.Timestamp{Time: time.Date(2021, time.May, 27, 11, 49, 05, 0, time.UTC)},
		}

		tm.products.EXPECT().Get(*product.ID).Times(1).Return(product, nil)

		config.Args = append(config.Args, *product.ID)

		err := runProductsGet(config)
		assert.NoError(t, err)
	})
}
