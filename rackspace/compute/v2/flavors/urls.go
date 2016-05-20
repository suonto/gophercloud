package flavors

import (
	"github.com/suonto/gophercloud"
)

func getURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("flavors", id)
}
