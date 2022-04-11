package commands

import (
	"testing"
	"time"

	"github.com/gocancel/gocancel-go"
	"github.com/stretchr/testify/assert"
)

func TestOrganizationsCommand(t *testing.T) {
	cmd := newOrganizationsCmd()
	assert.NotNil(t, cmd)
	assertCommandNames(t, cmd, "list", "get", "list-products", "get-product")
}

func TestOrganizationsList(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		organizations := []*gocancel.Organization{{
			ID:                gocancel.String("f38c8fab-0fa6-40b6-bb0c-6b3dfa2fec05"),
			Name:              gocancel.String("ACME"),
			Slug:              gocancel.String("acme"),
			Email:             gocancel.String("contact@acme.com"),
			URL:               gocancel.String("https://acme.com"),
			Phone:             gocancel.String("517-234-9141"),
			Fax:               gocancel.String("745-756-0818"),
			CategoryID:        gocancel.String("7df93680-4ce2-4da4-bdb1-d5a667695fd5"),
			RequiresConsent:   gocancel.Bool(true),
			RequiresProofOfID: gocancel.Bool(true),
			Metadata:          &gocancel.AccountMetadata{"foo": "bar"},
			CreatedAt:         &gocancel.Timestamp{Time: time.Date(2021, time.May, 27, 11, 49, 05, 0, time.UTC)},
			UpdatedAt:         &gocancel.Timestamp{Time: time.Date(2021, time.May, 27, 11, 49, 05, 0, time.UTC)},
		}}

		tm.organizations.EXPECT().List().Times(1).Return(organizations, nil)

		err := runOrganizationsList(config)
		assert.NoError(t, err)
	})
}

func TestOrganizationsGet(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		organization := &gocancel.Organization{
			ID:                gocancel.String("f38c8fab-0fa6-40b6-bb0c-6b3dfa2fec05"),
			Name:              gocancel.String("ACME"),
			Slug:              gocancel.String("acme"),
			Email:             gocancel.String("contact@acme.com"),
			URL:               gocancel.String("https://acme.com"),
			Phone:             gocancel.String("517-234-9141"),
			Fax:               gocancel.String("745-756-0818"),
			CategoryID:        gocancel.String("7df93680-4ce2-4da4-bdb1-d5a667695fd5"),
			RequiresConsent:   gocancel.Bool(true),
			RequiresProofOfID: gocancel.Bool(true),
			Metadata:          &gocancel.AccountMetadata{"foo": "bar"},
			CreatedAt:         &gocancel.Timestamp{Time: time.Date(2021, time.May, 27, 11, 49, 05, 0, time.UTC)},
			UpdatedAt:         &gocancel.Timestamp{Time: time.Date(2021, time.May, 27, 11, 49, 05, 0, time.UTC)},
		}

		tm.organizations.EXPECT().Get(*organization.ID).Times(1).Return(organization, nil)

		config.Args = append(config.Args, *organization.ID)

		err := runOrganizationsGet(config)
		assert.NoError(t, err)
	})
}

func TestOrganizationsListProducts(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		product := []*gocancel.Product{{
			ID:                gocancel.String("ad41acf6-9f3c-45cd-9438-134d6ee7831c"),
			Name:              gocancel.String("ACME"),
			Slug:              gocancel.String("acme"),
			Email:             gocancel.String("contact@acme.com"),
			URL:               gocancel.String("https://acme.com"),
			Phone:             gocancel.String("517-234-9141"),
			Fax:               gocancel.String("745-756-0818"),
			OrganizationID:    gocancel.String("f38c8fab-0fa6-40b6-bb0c-6b3dfa2fec05"),
			RequiresConsent:   gocancel.Bool(true),
			RequiresProofOfID: gocancel.Bool(true),
			Metadata:          &gocancel.AccountMetadata{"foo": "bar"},
			CreatedAt:         &gocancel.Timestamp{Time: time.Date(2021, time.May, 27, 11, 49, 05, 0, time.UTC)},
			UpdatedAt:         &gocancel.Timestamp{Time: time.Date(2021, time.May, 27, 11, 49, 05, 0, time.UTC)},
		}}

		tm.organizations.EXPECT().ListProducts(*product[0].OrganizationID).Times(1).Return(product, nil)

		config.Args = append(config.Args, *product[0].OrganizationID)

		err := runOrganizationsListProducts(config)
		assert.NoError(t, err)
	})
}

func TestOrganizationsGetProduct(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		product := &gocancel.Product{
			ID:                gocancel.String("ad41acf6-9f3c-45cd-9438-134d6ee7831c"),
			Name:              gocancel.String("ACME"),
			Slug:              gocancel.String("acme"),
			Email:             gocancel.String("contact@acme.com"),
			URL:               gocancel.String("https://acme.com"),
			Phone:             gocancel.String("517-234-9141"),
			Fax:               gocancel.String("745-756-0818"),
			OrganizationID:    gocancel.String("f38c8fab-0fa6-40b6-bb0c-6b3dfa2fec05"),
			RequiresConsent:   gocancel.Bool(true),
			RequiresProofOfID: gocancel.Bool(true),
			Metadata:          &gocancel.AccountMetadata{"foo": "bar"},
			CreatedAt:         &gocancel.Timestamp{Time: time.Date(2021, time.May, 27, 11, 49, 05, 0, time.UTC)},
			UpdatedAt:         &gocancel.Timestamp{Time: time.Date(2021, time.May, 27, 11, 49, 05, 0, time.UTC)},
		}

		tm.organizations.EXPECT().GetProduct(*product.OrganizationID, *product.ID).Times(1).Return(product, nil)

		config.Args = append(config.Args, *product.OrganizationID, *product.ID)

		err := runOrganizationsGetProduct(config)
		assert.NoError(t, err)
	})
}
