// +build acceptance

package rackspace

import (
	"testing"

	"github.com/suonto/gophercloud/acceptance/tools"
	"github.com/suonto/gophercloud/rackspace"
	th "github.com/suonto/gophercloud/testhelper"
)

func TestAuthenticatedClient(t *testing.T) {
	// Obtain credentials from the environment.
	ao, err := rackspace.AuthOptionsFromEnv()
	th.AssertNoErr(t, err)

	client, err := rackspace.AuthenticatedClient(tools.OnlyRS(ao))
	if err != nil {
		t.Fatalf("Unable to authenticate: %v", err)
	}

	if client.TokenID == "" {
		t.Errorf("No token ID assigned to the client")
	}

	t.Logf("Client successfully acquired a token: %v", client.TokenID)
}
