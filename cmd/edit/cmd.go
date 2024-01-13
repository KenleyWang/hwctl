package edit

import (
	"github.com/spf13/cobra"
	"hwctl/cmd/lts/log"
	"hwctl/cmd/lts/logStream"
)

//func init() {
//	client := InitClient("a", ",", "")
//}

var EditCmd = &cobra.Command{
	Use:          "edit",
	Short:        "编辑",
	SilenceUsage: true,
	Long:         `编辑`,
	//Example:      `awdw`,
	//PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	EditCmd.AddCommand(logStream.LogStreamCmd)
	EditCmd.AddCommand(log.LogCmd)
}
