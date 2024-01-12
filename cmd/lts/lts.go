package lts

import (
	"github.com/spf13/cobra"
	"hwctl/cmd/lts/cmd/log"
	"hwctl/cmd/lts/cmd/logGroup"
	"hwctl/cmd/lts/cmd/logStream"
)

//func init() {
//	client := InitClient("a", ",", "")
//}

var LtsCmd = &cobra.Command{
	Use:          "lts",
	Short:        "lts",
	SilenceUsage: true,
	Long:         `lts`,
	//Example:      `awdw`,
	//PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	//ltsCmd.SetHelpTemplate("daw")
	LtsCmd.AddCommand(logGroup.LogGroupCmd)
	LtsCmd.AddCommand(logStream.LogStreamCmd)
	LtsCmd.AddCommand(log.LogCmd)
}
