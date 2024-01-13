package log

import (
	"encoding/json"
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	ltsv2 "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/lts/v2"
	ltsModel "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/lts/v2/model"
	ltsRegion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/lts/v2/region"
	"hwctl/config"
	"hwctl/lib"
	"regexp"
	"strconv"
)

func getLogs(c *config.Config) {
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
	request := &ltsModel.ListLogsRequest{}
	//var listLabelsbody = map[string]string{
	//	"hostName": "ecs-kwxtest",
	//}
	limitQueryLtsLogParams := int32(500)
	keywordsQueryLtsLogParams := "methodNameDes"
	isCountQueryLtsLogParams := true
	request.LogGroupId = "526deca5-89d2-469e-a309-16dc90d69848"
	request.LogStreamId = "094183ca-1a34-4674-8417-d601e98fe59e"
	request.Body = &ltsModel.QueryLtsLogParams{
		Limit:    &limitQueryLtsLogParams,
		Keywords: &keywordsQueryLtsLogParams,
		IsCount:  &isCountQueryLtsLogParams,
		//Labels:    listLabelsbody,
		EndTime:   "1705035600000",
		StartTime: "1705032000000",
	}
	response, err := client.ListLogs(request)
	var mapR []map[string]string
	if err == nil {
		for _, v := range *response.Logs {
			logStruct, err := formatLogToStruct(*v.Content)
			if err != nil {
				panic(err)
			}
			//fmt.Printf("%+v\n", logStruct.Ip)
			//fmt.Printf("%+v\n", logStruct.MethodUrl)
			//fmt.Printf("%+v\n", logStruct.MemberId)
			//fmt.Printf("----------------------------")
			mapR = append(mapR, map[string]string{
				"ip":        logStruct.Ip,
				"MethodUrl": logStruct.MethodUrl,
				"memberID":  strconv.Itoa(logStruct.MemberId),
			})
		}
		err = lib.MapsToCSV(mapR, "./output.csv")
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println(err)
	}
}

type Log struct {
	Timestamp    string `json:"timestamp"`
	RequestId    string `json:"request_id"`
	Thread       string `json:"thread"`
	Level        string `json:"level"`
	Logger       string `json:"logger"`
	Method       string `json:"method"`
	MethodUrl    string `json:"methodUrl"`
	HighLightTag string `json:"<HighLightTag>methodNameDes</HighLightTag>"`
	Ip           string `json:"ip"`
	MemberId     int    `json:"memberId"`
}

func formatLogToStruct(logText string) (*Log, error) {
	// 定义日志文本的正则表达式
	logRegex := regexp.MustCompile(`(\d+:\d+:\d+\.\d+) \[(.*)] \[(.*)] (\w+) {2}([a-zA-Z.]+) - \[(.*)] - (.*)`)

	// 使用正则表达式提取信息
	matches := logRegex.FindStringSubmatch(logText)
	if len(matches) == 8 {

		// 构建Log对象
		log := &Log{
			Timestamp: matches[1],
			RequestId: matches[2],
			Thread:    matches[3],
			Level:     matches[4],
			Logger:    matches[5],
			Method:    matches[6],
		}

		// 解析JSON部分
		err := json.Unmarshal([]byte(matches[7]), log)
		if err != nil {
			return nil, err
		}

		return log, nil
	}

	return nil, fmt.Errorf("无法解析日志格式")
}
