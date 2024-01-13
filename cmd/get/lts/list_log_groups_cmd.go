package lts

import (
	ltsModel "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/lts/v2/model"
	"github.com/spf13/cobra"
	"hwctl/config"
	"hwctl/services/lts"
)

var listLogGroupsCmd = &cobra.Command{
	Use:          "loggroup",
	Short:        "获取华为云LTS(云日志服务)的日志组",
	SilenceUsage: true,
	Long:         `获取华为云LTS(云日志服务)的日志组`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c := cmd.Context().Value("c")
		if c, ok := c.(*config.Config); ok {
			var rsp *ltsModel.ListLogGroupsResponse
			ltsClient := lts.NewLTSClient(c)
			rsp, err := ltsClient.ListLogGroups()
			if err != nil {
				return err
			}
			println(rsp)
		}
		return nil
	},
}
