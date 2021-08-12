package block_storage

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumes"
	"github.com/gophercloud/gophercloud/pagination"
	"openstack-client/session"
	"openstack-client/utils"
)

/****************************
	Block Storage (Cinder)
*****************************/
func GetVolumeList(session session.OpenStackSession) {
	// 볼륨 목록 조회
	opts := gophercloud.EndpointOpts{Region: "RegionOne"}
	client, err := openstack.NewBlockStorageV3(session.Provider, opts)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Volume List")
	volumes.List(client, nil).EachPage(func (page pagination.Page) (bool, error) {
		volumeList, err := volumes.ExtractVolumes(page)
		if err != nil {
			return false, err
		}

		utils.PrintJson(volumeList)
		return true, nil;
	})
}
