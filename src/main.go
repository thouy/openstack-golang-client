package main

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumes"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/networks"
	"github.com/gophercloud/gophercloud/pagination"
)



type OpenStackSession struct {
	identityEndpoint string
	username string
	password string
	domainId string
	authOpts gophercloud.AuthOptions
	provider *gophercloud.ProviderClient
}

var session OpenStackSession
func initOpenstackSession(param map[string]string) {
	//session = OpenStackSession{}

	session.authOpts = gophercloud.AuthOptions{
		IdentityEndpoint: param["identityEndpoint"],
		Username: param["username"],
		Password: param["password"],
		DomainID: param["domainid"],
	}

	provider, err := openstack.AuthenticatedClient(session.authOpts)
	if err != nil {
		fmt.Println(err)
	}
	session.provider = provider
}

func getVolumeList(session OpenStackSession) {
	// 볼륨 목록 조회
	opts := gophercloud.EndpointOpts{Region: "RegionOne"}
	client, err := openstack.NewBlockStorageV3(session.provider, opts)
	if err != nil {
		fmt.Println(err)
	}

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
}

func getServerList(session OpenStackSession) {

	// 서버 목록 조회
	opts := gophercloud.EndpointOpts{Region: "RegionOne"}
	computeClient, computeErr := openstack.NewComputeV2(session.provider, opts)
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
}

func getNetworkList(session OpenStackSession) {

	// 네트워크 목록 조회
	opts := gophercloud.EndpointOpts{Region: "RegionOne"}
	networkClient, networkErr := openstack.NewNetworkV2(session.provider, opts)
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

func main() {
	params := make(map[string]string, 0)
	params["identityEndpoint"] = "http://203.255.255.100:5000"
	params["username"] = "paas-ta-monitoring"
	params["password"] = "paas-ta-monitoring!@#"
	params["domainid"] = "default"

	initOpenstackSession(params)

	getVolumeList(session)
	getServerList(session)
	getNetworkList(session)
}
