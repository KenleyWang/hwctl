package lts

import (
	"fmt"
	ltsModel "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/lts/v2/model"
	"github.com/spf13/cobra"
	"hwctl/config"
	"hwctl/pkg/printf"
	"hwctl/services/lts"
	"time"
)

var listLogsCmd = &cobra.Command{
	Use:          "log",
	Short:        "根据日志组和日志流获取华为云LTS(云日志服务)的日志",
	SilenceUsage: true,
	Long:         `获取华为云LTS(云日志服务)某个日志流下的日志`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c := cmd.Context().Value("c")
		if c, ok := c.(*config.Config); ok {
			var rsp *ltsModel.ListLogsResponse
			ltsClient := lts.NewLTSClient(c)
			logGroupID, err := cmd.Flags().GetString("log-group-id")
			if err != nil {
				return err
			}

			logStreamID, err := cmd.Flags().GetString("log-stream-id")
			if err != nil {
				return err
			}
			//cmd.Flags().GetString("config")
			//cmd.Flags().GetString("config")
			rsp, err = ltsClient.ListLogs(logGroupID, logStreamID, time.Now(), time.Now(), 1, "", true, map[string]string{})
			if err != nil {
				return err
			}
			fmt.Println(rsp.Logs)
			fmt.Println(rsp)
			printf.PrintSlice(rsp.Logs)
		}
		return nil
	},
}

func init() {
	listLogsCmd.Flags().StringP("log-group-id", "", "", "日志组ID")
	listLogsCmd.Flags().StringP("log-stream-id", "", "", "日志流ID")
	listLogsCmd.Flags().Int32P("limit", "", 50, "表示每次查询的日志条数,0<=limit<=5000")
	listLogsCmd.Flags().BoolP("is-count", "", true, "日志条数统计")
	listLogsCmd.Flags().StringP("keywords", "k", "", "支持关键词精确搜索。关键词指相邻两个分词符之间的单词,例：error")
	//listLogsCmd.Flags().Int32P("limit", "", 50, "0<=limit<=500")
}
