package client

import (
	"context"
	"fmt"
	"net/url"

	"github.com/gocancel/gocancel-go"
	"github.com/spf13/viper"
	"golang.org/x/oauth2/clientcredentials"
)

// Create returns a gocancel Client.
func Create(clientID string, clientSecret string) (*gocancel.Client, error) {
	if clientID == "" {
		return nil, fmt.Errorf("client identifier is required. (hint: set 'GOCANCEL_CLIENT_ID'")
	}

	if clientSecret == "" {
		return nil, fmt.Errorf("client secret is required. (hint: set 'GOCANCEL_CLIENT_SECRET'")
	}

	ctx := context.Background()

	base, _ := url.Parse(gocancel.Endpoint.TokenURL)
	token, _ := url.Parse("/oauth/token")

	apiURL := viper.GetString("api-url")
	if apiURL != "" {
		base, _ = url.Parse(apiURL + "/")
	}

	conf := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{"read:categories", "read:letters", "write:letters"},
		TokenURL:     base.ResolveReference(token).String(),
	}
	tc := conf.Client(ctx)

	return gocancel.New(tc, gocancel.SetBaseURL(base.String()))
}
