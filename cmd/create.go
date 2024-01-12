package cmd

import (
	"github.com/spf13/cobra"
	lts "hwctl/cmd/lts"
)

var createCmd = &cobra.Command{
	Use:          "create",
	Short:        "create",
	SilenceUsage: true,
	Long:         `create`,
	//PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	createCmd.AddCommand(lts.LtsCmd)
}
