package utils

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud/pagination"
	"reflect"
)

/*************
	 Utils
**************/
func PrintJson(p interface{}) {
	doc, err := json.MarshalIndent(p, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(doc))
}

func PagerToMap(pager pagination.Pager) interface{} {
	allPages, err := pager.AllPages()
	if err != nil {
		fmt.Println(err)
	}
	pageBody := allPages.GetBody()

	fmt.Printf("pageBody type is %s\n", reflect.TypeOf(pageBody))

	return pageBody
}
