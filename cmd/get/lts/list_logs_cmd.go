package lts

import (
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
			limit, err := cmd.Flags().GetInt32("limit")
			if err != nil {
				return err
			}
			keywords, err := cmd.Flags().GetString("keywords")
			if err != nil {
				return err
			}

			min, err := cmd.Flags().GetInt32("min")
			if err != nil {
				return err
			}
			currentTime := time.Now()

			// 计算 n 分钟前的时间
			beforeMin := currentTime.Add(-time.Duration(min) * time.Minute)
			rsp, err = ltsClient.ListLogs(logGroupID, logStreamID, beforeMin, time.Now(), limit, keywords, true, map[string]string{})
			if err != nil {
				return err
			}
			//fmt.Println(rsp.Logs)
			//fmt.Println(rsp)
			printf.PrintSlice(rsp.Logs)
		}
		return nil
	},
}

func init() {
	listLogsCmd.Flags().StringP("log-group-id", "", "", "日志组ID")
	listLogsCmd.Flags().StringP("log-stream-id", "", "", "日志流ID")
	listLogsCmd.Flags().Int32P("limit", "", 50, "表示每次查询的日志条数,0<=limit<=5000")
	listLogsCmd.Flags().BoolP("is-count", "", true, "是否统计日志条数")
	listLogsCmd.Flags().StringP("keywords", "k", "", "支持关键词精确搜索。关键词指相邻两个分词符之间的单词,例：error")
	listLogsCmd.Flags().Int32P("min", "", 60, "查找在n分钟内的日志,0<=min<=24*60")
	listLogsCmd.Flags().Int32P("hour", "", 1, "查找在n小时内的日志,0<=hour<=24")
	err := listLogsCmd.MarkFlagRequired("log-group-id")
	if err != nil {
		panic(err)
	}
	err = listLogsCmd.MarkFlagRequired("log-stream-id")
	if err != nil {
		panic(err)
	}
}
