package log_group

import (
	ltsv2 "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/lts/v2"
	ltsModel "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/lts/v2/model"
	ltsRegion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/lts/v2/region"
	"hwctl/config"
)

func ListLogGroup(c *config.Config) (*ltsModel.ListLogGroupsResponse, error) {

	//创建服务客户端
	client := ltsv2.NewLtsClient(
		ltsv2.LtsClientBuilder().
			WithRegion(ltsRegion.ValueOf(c.Region)).
			WithCredential(c.GenerateCredential()).
			Build())
	request := &ltsModel.ListLogGroupsRequest{}
	response, err := client.ListLogGroups(request)
	return response, err
}
