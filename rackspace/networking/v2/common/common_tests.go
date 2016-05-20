package common

import (
	"github.com/suonto/gophercloud"
	"github.com/suonto/gophercloud/testhelper/client"
)

const TokenID = client.TokenID

func ServiceClient() *gophercloud.ServiceClient {
	return client.ServiceClient()
}
