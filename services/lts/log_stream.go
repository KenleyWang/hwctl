package lts

import ltsModel "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/lts/v2/model"

type (
	LogStream interface {
		ListLogStreams() (*ltsModel.ListLogStreamsResponse, error)
	}
)

func (c *ltsClient) ListLogStreams() (*ltsModel.ListLogStreamsResponse, error) {
	return nil, nil
}
