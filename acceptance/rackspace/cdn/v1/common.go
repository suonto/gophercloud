// +build acceptance

package v1

import (
	"testing"

	"github.com/suonto/gophercloud"
	"github.com/suonto/gophercloud/rackspace"
	th "github.com/suonto/gophercloud/testhelper"
)

func newClient(t *testing.T) *gophercloud.ServiceClient {
	ao, err := rackspace.AuthOptionsFromEnv()
	th.AssertNoErr(t, err)

	client, err := rackspace.AuthenticatedClient(ao)
	th.AssertNoErr(t, err)

	c, err := rackspace.NewCDNV1(client, gophercloud.EndpointOpts{})
	th.AssertNoErr(t, err)
	return c
}
