package commands

import (
	"io/ioutil"
	"sort"
	"testing"

	"github.com/gocancel/gocancel-cli/client"
	"github.com/gocancel/gocancel-cli/client/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func assertCommandNames(t *testing.T, cmd *Command, expected ...string) {
	var names []string

	for _, c := range cmd.Commands() {
		names = append(names, c.Name())
		if c.Name() == "list" {
			assert.Contains(t, c.Aliases, "ls", "Missing 'ls' alias for 'list' command.")
			assert.NotNil(t, c.Flags().Lookup("format"), "Missing 'format' flag for 'list' command.")
		}

		if c.Name() == "get" {
			assert.NotNil(t, c.Flags().Lookup("format"), "Missing 'format' flag for 'get' command.")
		}
	}

	sort.Strings(expected)
	sort.Strings(names)
	assert.Equal(t, expected, names)
}

type testFn func(c *CmdConfig, tm *tcMocks)

type tcMocks struct {
	categories    *mocks.MockCategoriesService
	letters       *mocks.MockLettersService
	organizations *mocks.MockOrganizationsService
	products      *mocks.MockProductsService
}

func withTestClient(t *testing.T, tFn testFn) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tm := &tcMocks{
		categories:    mocks.NewMockCategoriesService(ctrl),
		letters:       mocks.NewMockLettersService(ctrl),
		organizations: mocks.NewMockOrganizationsService(ctrl),
		products:      mocks.NewMockProductsService(ctrl),
	}

	config := &CmdConfig{
		NS:  "test",
		Out: ioutil.Discard,

		// can stub this out, since the return is dictated by the mocks.
		initServices: func(c *CmdConfig) error { return nil },

		Categories:    func() client.CategoriesService { return tm.categories },
		Letters:       func() client.LettersService { return tm.letters },
		Organizations: func() client.OrganizationsService { return tm.organizations },
		Products:      func() client.ProductsService { return tm.products },
	}

	tFn(config, tm)
}
