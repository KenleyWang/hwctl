package del

import (
	"github.com/spf13/cobra"
)

//func init() {
//	client := InitClient("a", ",", "")
//}

var DeleteCmd = &cobra.Command{
	Use:          "delete",
	Short:        "删除",
	SilenceUsage: true,
	Long:         `删除`,
	//Example:      `awdw`,
	//PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	//DeleteCmd.AddCommand(logStream.LogStreamCmd)
	//DeleteCmd.AddCommand(log.LogCmd)
}
