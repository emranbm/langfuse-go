package api

import (
	"context"
	"encoding/base64"
	"net/http"
	"os"

	"github.com/henomis/restclientgo"
)

const (
	langfuseDefaultEndpoint = "https://cloud.langfuse.com"
)

type Client struct {
	restClient *restclientgo.RestClient
}

// New creates a new Langfuse client with the environment variables:
// LANGFUSE_HOST, LANGFUSE_PUBLIC_KEY, and LANGFUSE_SECRET_KEY.
func New() *Client {
	langfuseHost := os.Getenv("LANGFUSE_HOST")
	publicKey := os.Getenv("LANGFUSE_PUBLIC_KEY")
	secretKey := os.Getenv("LANGFUSE_SECRET_KEY")
	return NewWithParams(langfuseHost, publicKey, secretKey)

}

// NewWithParams creates a new Langfuse client with the specified endpoint, public key, and secret key.
// If the endpoint is empty, it defaults to the Langfuse cloud endpoint.
func NewWithParams(endpoint, publicKey, secretKey string) *Client {
	if endpoint == "" {
		endpoint = langfuseDefaultEndpoint
	}
	restClient := restclientgo.New(endpoint)
	restClient.SetRequestModifier(func(req *http.Request) *http.Request {
		req.Header.Set("Authorization", basicAuth(publicKey, secretKey))
		return req
	})

	return &Client{
		restClient: restClient,
	}
}

func (c *Client) Ingestion(ctx context.Context, req *Ingestion, res *IngestionResponse) error {
	return c.restClient.Post(ctx, req, res)
}

func basicAuth(publicKey, secretKey string) string {
	auth := publicKey + ":" + secretKey
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}
