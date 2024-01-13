package get

import (
	"github.com/spf13/cobra"
	"hwctl/cmd/get/lts"
)

//func init() {
//	client := InitClient("a", ",", "")
//}

var GetCmd = &cobra.Command{
	Use:          "get",
	Short:        "获取华为云产品中的一项资源或服务",
	SilenceUsage: true,
	Long:         `获取华为云产品中的一项资源或服务`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	//GetCmd.AddCommand(logStream.LogStreamCmd)
	GetCmd.AddCommand(lts.LtsCmd)
}
