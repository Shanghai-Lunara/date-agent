package date_agent

import (
	"k8s.io/klog"
)

type Client struct {
	hostname     string
	registerAddr string
}

// NewClient returns the pointer of the Client structure
func NewClient(addr string) *Client {
	hostname, err := GetHostName()
	if err != nil {
		klog.Fatal(err)
	}
	c := &Client{
		hostname:     hostname,
		registerAddr: addr,
	}
	return c
}

func (c *Client) Loop() {

}
