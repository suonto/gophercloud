package buildinfo

import (
	"testing"

	os "github.com/suonto/gophercloud/openstack/orchestration/v1/buildinfo"
	th "github.com/suonto/gophercloud/testhelper"
	fake "github.com/suonto/gophercloud/testhelper/client"
)

func TestGetTemplate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	os.HandleGetSuccessfully(t, os.GetOutput)

	actual, err := Get(fake.ServiceClient()).Extract()
	th.AssertNoErr(t, err)

	expected := os.GetExpected
	th.AssertDeepEquals(t, expected, actual)
}
