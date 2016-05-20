package apiversions

import "github.com/suonto/gophercloud"

func apiVersionsURL(c *gophercloud.ServiceClient) string {
	return c.Endpoint
}
