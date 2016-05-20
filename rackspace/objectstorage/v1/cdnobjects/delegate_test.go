package cdnobjects

import (
	"testing"

	os "github.com/suonto/gophercloud/openstack/objectstorage/v1/objects"
	th "github.com/suonto/gophercloud/testhelper"
	fake "github.com/suonto/gophercloud/testhelper/client"
)

func TestDeleteCDNObject(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	os.HandleDeleteObjectSuccessfully(t)

	res := Delete(fake.ServiceClient(), "testContainer", "testObject", nil)
	th.AssertNoErr(t, res.Err)

}
