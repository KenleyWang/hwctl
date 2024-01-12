package cmd

import (
	"github.com/spf13/cobra"
	lts "hwctl/cmd/lts"
)

var deleteCmd = &cobra.Command{
	Use:          "delete",
	Short:        "delete",
	SilenceUsage: true,
	Long:         `del`,
	//PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	deleteCmd.AddCommand(lts.LtsCmd)
}
