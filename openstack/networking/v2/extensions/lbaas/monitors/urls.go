package monitors

import "github.com/suonto/gophercloud"

const (
	rootPath     = "lb"
	resourcePath = "health_monitors"
)

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(rootPath, resourcePath)
}

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id)
}
