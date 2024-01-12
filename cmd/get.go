package cmd

import (
	"github.com/spf13/cobra"
	lts "hwctl/cmd/lts"
)

var getCmd = &cobra.Command{
	Use:          "get",
	Short:        "get",
	SilenceUsage: true,
	Long:         `get`,
	//PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	getCmd.AddCommand(lts.LtsCmd)
}
