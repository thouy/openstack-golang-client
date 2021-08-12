package compute

import (
	"fmt"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/flavors"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/pagination"

	"openstack-client/session"
	"openstack-client/utils"
)

/********************
	Compute (Nova)
*********************/

func GetServerList(session session.OpenStackSession) {
	// 서버 목록 조회
	opts := gophercloud.EndpointOpts{
		Region: "RegionOne",
	}
	computeClient, computeErr := openstack.NewComputeV2(session.Provider, opts)
	if computeErr != nil {
		fmt.Println(computeErr)
	}

	fmt.Println("Server List")
	servers.List(computeClient, nil).EachPage(func (page pagination.Page) (bool, error) {
		serverList, err := servers.ExtractServers(page)
		if err != nil {
			return false, err
		}

		utils.PrintJson(serverList)
		return true, nil;
	})
}

func GetServerDetail(session session.OpenStackSession, serverId string) {
	opts := gophercloud.EndpointOpts{
		Region: "RegionOne",
	}

	serverClient, err := openstack.NewComputeV2(session.Provider, opts)
	if err != nil {
		fmt.Println(err)
	}

	result := servers.Get(serverClient, serverId)

	utils.PrintJson(result.Body)
}

func GetFlavor(session session.OpenStackSession, flavorId string) {
	// 네트워크 목록 조회
	opts := gophercloud.EndpointOpts{Region: "RegionOne"}
	flavorClient, err := openstack.NewComputeV2(session.Provider, opts)
	if err != nil {
		fmt.Println(err)
	}

	result := flavors.Get(flavorClient, flavorId)

	utils.PrintJson(result.Body)
}