package main

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/networks"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumes"
	"github.com/gophercloud/gophercloud/pagination"
)


func main() {
	authOpts := gophercloud.AuthOptions{
		IdentityEndpoint: "http://203.255.255.100:5000",
		Username: "paas-ta-monitoring",
		Password: "paas-ta-monitoring!@#",
		DomainID: "default",
	}

	provider, err := openstack.AuthenticatedClient(authOpts)
	if err != nil {
		fmt.Print(err)
	}

	// 볼륨 목록 조회
	opts := gophercloud.EndpointOpts{Region: "RegionOne"}
	client, err := openstack.NewBlockStorageV3(provider, opts)

	fmt.Println("Volume List")
	volumes.List(client, nil).EachPage(func (page pagination.Page) (bool, error) {
		s, err := volumes.ExtractVolumes(page)
		if err != nil {
			return false, err
		}
		doc, err := json.MarshalIndent(s, "", " ")
		fmt.Println(string(doc))

		return true, nil;
	})

	// 서버 목록 조회
	computeClient, computeErr := openstack.NewComputeV2(provider, opts)
	if computeErr != nil {
		fmt.Println(computeErr)
	}

	fmt.Println("Server List")
	servers.List(computeClient, nil).EachPage(func (page pagination.Page) (bool, error) {
		serverList, err := servers.ExtractServers(page)
		if err != nil {
			return false, err
		}

		doc, err := json.MarshalIndent(serverList, "", " ")
		fmt.Println(string(doc))

		return true, nil;
	})

	// 네트워크 목록 조회
	networkClient, networkErr := openstack.NewNetworkV2(provider, opts)
	if networkErr != nil {
		fmt.Println(networkErr)
	}


	fmt.Println("Network List")
	networks.List(networkClient, nil).EachPage(func (page pagination.Page) (bool, error) {
		networkList, err := networks.ExtractNetworks(page)
		if err != nil {
			return false, err
		}

		doc, err := json.MarshalIndent(networkList, "", " ")
		fmt.Println(string(doc))

		return true, nil;
	})
}
