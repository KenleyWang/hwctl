package lts

import (
	"fmt"
	ltsModel "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/lts/v2/model"
	"time"
)

type (
	Log interface {
		ListLogs() (*ltsModel.ListLogsResponse, error)
	}
)

func (c *ltsClient) ListLogs(LogGroupId string, LogStreamId string, startTime time.Time, endTime time.Time, limit int32, keywords string, isCount bool, listLabelled map[string]string) (*ltsModel.ListLogsResponse, error) {
	request := &ltsModel.ListLogsRequest{}
	request.LogGroupId = LogGroupId
	request.LogStreamId = LogStreamId
	request.Body = &ltsModel.QueryLtsLogParams{
		Limit:     &limit,
		Keywords:  &keywords,
		IsCount:   &isCount,
		Labels:    listLabelled,
		EndTime:   fmt.Sprintf("%d", endTime.UnixNano()/int64(time.Millisecond)),
		StartTime: fmt.Sprintf("%d", startTime.UnixNano()/int64(time.Millisecond)),
	}
	response, err := c.client.ListLogs(request)
	return response, err
}
