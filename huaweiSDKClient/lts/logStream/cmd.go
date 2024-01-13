package logStream

import "github.com/spf13/cobra"

var LogStreamCmd = &cobra.Command{
	Use:          "logstream",
	Short:        "获取日志流信息",
	SilenceUsage: true,
	Long:         `logstream----`,
	//Example:      `awdw`,
	//PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
