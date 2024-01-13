package lts

import ltsModel "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/lts/v2/model"

type (
	LogGroup interface {
		ListLogGroups() (*ltsModel.ListLogGroupsResponse, error)
	}
)

func (c *ltsClient) ListLogGroups() (*ltsModel.ListLogGroupsResponse, error) {
	return nil, nil
}
