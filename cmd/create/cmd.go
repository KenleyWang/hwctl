package create

import (
	"github.com/spf13/cobra"
)

//func init() {
//	client := InitClient("a", ",", "")
//}

var CreateCmd = &cobra.Command{
	Use:          "create",
	Short:        "创建",
	SilenceUsage: true,
	Long:         `创建`,
	//Example:      `awdw`,
	//PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	//CreateCmd.AddCommand(logStream.LogStreamCmd)
	//CreateCmd.AddCommand(log.LogCmd)
}
