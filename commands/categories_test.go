package commands

import (
	"testing"
	"time"

	"github.com/gocancel/gocancel-go"
	"github.com/stretchr/testify/assert"
)

func TestCategoriesCommand(t *testing.T) {
	cmd := newCategoriesCmd()
	assert.NotNil(t, cmd)
	assertCommandNames(t, cmd, "list", "get")
}

func TestCategoriesList(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		categories := []*gocancel.Category{{
			ID:              gocancel.String("7df93680-4ce2-4da4-bdb1-d5a667695fd5"),
			Name:            gocancel.String("Foo"),
			Slug:            gocancel.String("foo"),
			RequiresConsent: gocancel.Bool(false),
			Metadata:        &gocancel.AccountMetadata{"foo": "bar"},
			CreatedAt:       &gocancel.Timestamp{Time: time.Date(2021, time.May, 27, 11, 49, 05, 0, time.UTC)},
			UpdatedAt:       &gocancel.Timestamp{Time: time.Date(2021, time.May, 27, 11, 49, 05, 0, time.UTC)},
		}}

		tm.categories.EXPECT().List().Times(1).Return(categories, nil)

		err := runCategoriesList(config)
		assert.NoError(t, err)
	})
}

func TestCategoriesGet(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		category := &gocancel.Category{
			ID:              gocancel.String("7df93680-4ce2-4da4-bdb1-d5a667695fd5"),
			Name:            gocancel.String("Foo"),
			Slug:            gocancel.String("foo"),
			RequiresConsent: gocancel.Bool(false),
			Metadata:        &gocancel.AccountMetadata{"foo": "bar"},
			CreatedAt:       &gocancel.Timestamp{Time: time.Date(2021, time.May, 27, 11, 49, 05, 0, time.UTC)},
			UpdatedAt:       &gocancel.Timestamp{Time: time.Date(2021, time.May, 27, 11, 49, 05, 0, time.UTC)},
		}

		tm.categories.EXPECT().Get(*category.ID).Times(1).Return(category, nil)

		config.Args = append(config.Args, *category.ID)

		err := runCategoriesGet(config)
		assert.NoError(t, err)
	})
}
