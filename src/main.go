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

	var volumeParams = make(map[string]string, 0)
	volumeList := block_storage.GetVolumeList(osSession, volumeParams)
	fmt.Printf("Volumes count : %d\n", len(volumeList))

	var networkParams = make(map[string]string, 0)
	networkList := network.GetNetworkList(osSession, networkParams)
	fmt.Printf("Floating IP count : %d\n", len(networkList))

	var params = make(map[string]string, 0)
	params["status"] = "DOWN"
	floatingipList := network.GetFloatingIps(osSession, params)
	fmt.Printf("Floating IP count : %d\n", len(floatingipList))

	securityGroupList := network.GetSecurityGroups(osSession,nil)
	fmt.Printf("Security group count : %d\n", len(securityGroupList))

	// 전체 서버 목록 조회
	serverParams := make(map[string]interface{}, 0)
	serverParams["allTenants"] = true
	serverParams["tenantId"] = "944774277a034e648ab72e318db3f976"   // paas-ta-monitoring 프로젝트
	compute.GetServerList(osSession, serverParams)
	
	// 서버 상세정보 조회
	compute.GetServerDetail(osSession, "9d3ab7fa-fcaf-4750-badb-1e39f785ccb2")
	
	// Flaver 조회
	compute.GetFlavor(osSession,"4")
	
	// Hypervisor 전체 통계 조회
	compute.GetHypervisorStatistics(osSession)

	// all01 정보 조회
	hypervisorParams := make(map[string]interface{}, 0)
	hypervisorParams["hypervisorHostname"] = "all01"
	//hypervisorParams["withServers"] = true
	compute.GetHypervisorList(osSession, hypervisorParams)

	compute.GetHypervisor(osSession, "8e4f7157-4b5b-4d90-a7c6-5fc098b31f32")
	compute.GetHypervisor(osSession, "1379818a-bc5d-4d5e-97c6-9b33620c9a71")

	// 프로젝트 목록 조회
	compute.GetProjectList(osSession)

}
