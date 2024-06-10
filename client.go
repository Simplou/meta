package meta

import (
	"bytes"
	"io"
	"net/http"

	"github.com/Simplou/goxios"
)

var BaseUrl = "https://graph.facebook.com"

type HttpClient interface {
	Get(string, *goxios.RequestOpts) (*http.Response, error)
	Post(string, *goxios.RequestOpts) (*http.Response, error)
	Put(string, *goxios.RequestOpts) (*http.Response, error)
	Patch(string, *goxios.RequestOpts) (*http.Response, error)
	Delete(string, *goxios.RequestOpts) (*http.Response, error)
	SetHeaders(headers ...goxios.Header) []goxios.Header
}

type Client struct {
	HttpClient
	ApiVersion  string
	BaseUrl     string
	AccessToken string
}

func (c Client) Buffer(buf []byte) *bytes.Buffer {
	return bytes.NewBuffer(buf)
}

func (c Client) Reader(buf []byte) io.Reader {
	return c.Buffer(buf)
}

func (c Client) Writer(buf []byte) io.Writer {
	return c.Buffer(buf)
}
