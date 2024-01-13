package lts

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	lts "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/lts/v2"
	ltsRegion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/lts/v2/region"
	"hwctl/config"
)

type (
	LTSClient interface {
		GetClient() *lts.LtsClient
	}
	ltsClient struct {
		Client *lts.LtsClient
	}
)

func (client *ltsClient) GetClient() *lts.LtsClient {
	return client.Client
}

func NewLTSClient(c *config.Config) LTSClient {
	return &ltsClient{
		Client: initClient(c),
	}

}

func initClient(c *config.Config) *lts.LtsClient {
	// 配置认证信息
	auth := basic.NewCredentialsBuilder().
		WithAk(c.AK).
		WithSk(c.SK).
		Build()
	//创建服务客户端
	return lts.NewLtsClient(
		lts.LtsClientBuilder().
			WithRegion(ltsRegion.ValueOf(c.Region)).
			WithCredential(auth).
			Build())
}
