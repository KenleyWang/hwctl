package lts

import ltsModel "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/lts/v2/model"

type (
	LogStream interface {
		ListLogStream(string) (*ltsModel.ListLogStreamResponse, error)
	}
)

func (c *ltsClient) ListLogStream(logGroupId string) (*ltsModel.ListLogStreamResponse, error) {
	request := &ltsModel.ListLogStreamRequest{
		LogGroupId: logGroupId,
	}
	response, err := c.client.ListLogStream(request)
	return response, err
}
