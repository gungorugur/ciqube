package httpclient

import "net/http"

//Client is wrapper of http.client
type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

//New returns wrapped http client
func New(client *http.Client) Client {
	return defaultClient{client: client}
}

type defaultClient struct {
	client *http.Client
}

func (c defaultClient) Do(req *http.Request) (*http.Response, error) {
	return c.client.Do(req)
}
