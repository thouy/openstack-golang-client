package main

import (
	"fmt"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"

	"openstack-client/block-storage"
	"openstack-client/compute"
	"openstack-client/network"
	"openstack-client/session"
)

var osSession session.OpenStackSession

func initOpenstackSession(params map[string]string) {

	osSession.AuthOpts = gophercloud.AuthOptions{
		IdentityEndpoint: params["identityEndpoint"],
		Username: params["username"],
		Password: params["password"],
		DomainID: params["domainid"],
	}

	provider, err := openstack.AuthenticatedClient(osSession.AuthOpts)
	if err != nil {
		fmt.Println(err)
	}
	osSession.Provider = provider
}

func main() {
	sessionParams := make(map[string]string, 0)
	sessionParams["identityEndpoint"] = "http://203.255.255.100:5000"
	sessionParams["username"] = "paas-ta-monitoring"
	sessionParams["password"] = "paas-ta-monitoring!@#"
	sessionParams["domainid"] = "default"

	initOpenstackSession(sessionParams)

	block_storage.GetVolumeList(osSession)

	network.GetNetworkList(osSession)

	var params = make(map[string]string, 0)
	params["status"] = "DOWN"
	floatingipList := network.GetFloatingIps(osSession, params)
	fmt.Printf("Floating IP count : %d\n", len(floatingipList))

	securityGroupList := network.GetSecurityGroups(osSession,nil)
	fmt.Printf("Security group count : %d\n", len(securityGroupList))

	compute.GetServerList(osSession)
	compute.GetServerDetail(osSession, "9d3ab7fa-fcaf-4750-badb-1e39f785ccb2")
	compute.GetFlavor(osSession,"4")
	compute.GetHypervisorStatistics(osSession)
}