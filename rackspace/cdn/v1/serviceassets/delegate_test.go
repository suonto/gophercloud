package serviceassets

import (
	"testing"

	os "github.com/suonto/gophercloud/openstack/cdn/v1/serviceassets"
	th "github.com/suonto/gophercloud/testhelper"
	fake "github.com/suonto/gophercloud/testhelper/client"
)

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	os.HandleDeleteCDNAssetSuccessfully(t)

	err := Delete(fake.ServiceClient(), "96737ae3-cfc1-4c72-be88-5d0e7cc9a3f0", nil).ExtractErr()
	th.AssertNoErr(t, err)
}
