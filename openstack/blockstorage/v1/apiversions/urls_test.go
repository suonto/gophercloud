package apiversions

import (
	"testing"

	"github.com/suonto/gophercloud"
	th "github.com/suonto/gophercloud/testhelper"
)

const endpoint = "http://localhost:57909/"

func endpointClient() *gophercloud.ServiceClient {
	return &gophercloud.ServiceClient{Endpoint: endpoint}
}

func TestGetURL(t *testing.T) {
	actual := getURL(endpointClient(), "v1")
	expected := endpoint + "v1/"
	th.AssertEquals(t, expected, actual)
}

func TestListURL(t *testing.T) {
	actual := listURL(endpointClient())
	expected := endpoint
	th.AssertEquals(t, expected, actual)
}
