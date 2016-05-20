package tokens

import (
	"testing"

	"github.com/suonto/gophercloud"
	os "github.com/suonto/gophercloud/openstack/identity/v2/tokens"
	th "github.com/suonto/gophercloud/testhelper"
	"github.com/suonto/gophercloud/testhelper/client"
)

func tokenPost(t *testing.T, options gophercloud.AuthOptions, requestJSON string) os.CreateResult {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	os.HandleTokenPost(t, requestJSON)

	return Create(client.ServiceClient(), WrapOptions(options))
}

func TestCreateTokenWithAPIKey(t *testing.T) {
	options := gophercloud.AuthOptions{
		Username: "me",
		APIKey:   "1234567890abcdef",
	}

	os.IsSuccessful(t, tokenPost(t, options, `
    {
      "auth": {
        "RAX-KSKEY:apiKeyCredentials": {
          "username": "me",
          "apiKey": "1234567890abcdef"
        }
      }
    }
  `))
}
