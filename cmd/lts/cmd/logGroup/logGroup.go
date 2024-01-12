package logGroup

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	ltsv2 "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/lts/v2"
	ltsModel "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/lts/v2/model"
	ltsRegion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/lts/v2/region"
	"hwctl/config"
)

func getLogGroup(c *config.Config) {
	auth := basic.NewCredentialsBuilder().
		WithAk(c.AK).
		WithSk(c.SK).
		Build()

	//创建服务客户端
	client := ltsv2.NewLtsClient(
		ltsv2.LtsClientBuilder().
			WithRegion(ltsRegion.ValueOf(c.Region)).
			WithCredential(auth).
			Build())
	request := &ltsModel.ListLogGroupsRequest{}
	response, _ := client.ListLogGroups(request)
	for i, value := range *response.LogGroups {
		fmt.Printf("Index %d: %d\n", i, value)
	}
	//request := &ltsModel.ListLogGroupsRequest{}
	//response, err := client.ListLogGroups(request)
	//if err == nil {
	//	fmt.Printf("%+v\n", response)
	//} else {
	//	fmt.Println(err)
	//}

}
