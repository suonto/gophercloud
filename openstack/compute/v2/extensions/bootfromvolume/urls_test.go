package bootfromvolume

import (
	"testing"

	th "github.com/suonto/gophercloud/testhelper"
	"github.com/suonto/gophercloud/testhelper/client"
)

func TestCreateURL(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	c := client.ServiceClient()

	th.CheckEquals(t, c.Endpoint+"os-volumes_boot", createURL(c))
}
