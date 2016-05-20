// +build acceptance

package v2

import (
	"os"
	"testing"

	"github.com/suonto/gophercloud"
	"github.com/suonto/gophercloud/acceptance/tools"
	"github.com/suonto/gophercloud/rackspace"
	"github.com/suonto/gophercloud/rackspace/identity/v2/tokens"
	th "github.com/suonto/gophercloud/testhelper"
)

func rackspaceAuthOptions(t *testing.T) gophercloud.AuthOptions {
	// Obtain credentials from the environment.
	options, err := rackspace.AuthOptionsFromEnv()
	th.AssertNoErr(t, err)
	options = tools.OnlyRS(options)

	if options.Username == "" {
		t.Fatal("Please provide a Rackspace username as RS_USERNAME.")
	}
	if options.APIKey == "" {
		t.Fatal("Please provide a Rackspace API key as RS_API_KEY.")
	}

	return options
}

func createClient(t *testing.T, auth bool) *gophercloud.ServiceClient {
	ao := rackspaceAuthOptions(t)

	provider, err := rackspace.NewClient(ao.IdentityEndpoint)
	th.AssertNoErr(t, err)

	if auth {
		err = rackspace.Authenticate(provider, ao)
		th.AssertNoErr(t, err)
	}

	return rackspace.NewIdentityV2(provider)
}

func TestTokenAuth(t *testing.T) {
	authedClient := createClient(t, true)
	token := authedClient.TokenID

	tenantID := os.Getenv("RS_TENANT_ID")
	if tenantID == "" {
		t.Skip("You must set RS_TENANT_ID environment variable to run this test")
	}

	authOpts := tokens.AuthOptions{}
	authOpts.TenantID = tenantID
	authOpts.Token = token

	_, err := tokens.Create(authedClient, authOpts).ExtractToken()
	th.AssertNoErr(t, err)
}
