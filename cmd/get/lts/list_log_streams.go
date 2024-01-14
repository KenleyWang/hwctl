package lts

import (
	ltsModel "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/lts/v2/model"
	"github.com/spf13/cobra"
	"hwctl/config"
	"hwctl/pkg/printf"
	"hwctl/services/lts"
)

var listLogStreamsCmd = &cobra.Command{
	Use:          "logstream",
	Short:        "根据日志组获取华为云LTS(云日志服务)的日志流",
	SilenceUsage: true,
	Long:         `根据日志组获取华为云LTS(云日志服务)的日志流`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c := cmd.Context().Value("c")
		if c, ok := c.(*config.Config); ok {
			var rsp *ltsModel.ListLogStreamResponse
			ltsClient := lts.NewLTSClient(c)
			logGroupID, err := cmd.Flags().GetString("log-group-id")
			if err != nil {
				return err
			}
			rsp, err = ltsClient.ListLogStream(logGroupID)
			if err != nil {
				return err
			}
			printf.PrintSlice(rsp.LogStreams)
		}
		return nil
	},
}

func init() {
	listLogStreamsCmd.Flags().StringP("log-group-id", "", "", "日志组ID")
}
