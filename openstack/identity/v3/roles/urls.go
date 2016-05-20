package roles

import "github.com/suonto/gophercloud"

func listAssignmentsURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("role_assignments")
}
