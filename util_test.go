package gophercloud

import (
	"testing"

	th "github.com/suonto/gophercloud/testhelper"
)

func TestWaitFor(t *testing.T) {
	err := WaitFor(5, func() (bool, error) {
		return true, nil
	})
	th.CheckNoErr(t, err)
}
