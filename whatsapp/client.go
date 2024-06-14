package whatsapp

import (
	"context"
	"fmt"

	"github.com/Simplou/goxios"
	"github.com/Simplou/meta"
)

type client struct {
	senderID string
	meta.Client
}

func New(ctx context.Context, apiVersion, accessToken, senderID string) *client {
	return &client{
		senderID: senderID,
		Client: meta.Client{
			HttpClient:  goxios.New(ctx),
			ApiVersion:  apiVersion,
			BaseUrl:     meta.BaseUrl,
			AccessToken: accessToken,
		},
	}
}

func (c *client) Endpoint(endpoint string) string {
	return fmt.Sprintf("%s/%s/%s", c.BaseUrl, c.ApiVersion, endpoint)
}

func (c *client) Headers(contentType string) []goxios.Header {
	return c.SetHeaders(goxios.Header{
		Key:   "Content-Type",
		Value: contentType,
	}, goxios.Header{
		Key:   "Authorization",
		Value: fmt.Sprintf("Bearer %s", c.AccessToken),
	})
}
