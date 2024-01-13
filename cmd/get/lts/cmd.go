package lts

import (
	"github.com/spf13/cobra"
)

var LtsCmd = &cobra.Command{
	Use:          "lts",
	Short:        "获取华为云LTS(云日志服务)的日志组",
	SilenceUsage: true,
	Long:         `获取华为云LTS(云日志服务)的日志组`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	LtsCmd.AddCommand(listLogGroupCmd)
}
