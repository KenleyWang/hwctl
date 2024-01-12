package logGroup

import (
	"github.com/spf13/cobra"
	"hwctl/config"
)

var LogGroupCmd = &cobra.Command{
	Use:          "loggroup",
	Short:        "loggroup",
	SilenceUsage: true,
	Long:         `loggroup`,
	//Example:      `awdw`,
	//PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: cmdRun,
}

func cmdRun(cmd *cobra.Command, args []string) {

	s, err := cmd.Flags().GetString("config")
	if err != nil {
		panic(err)
	}

	var c config.Config
	c.MustLoadConfig(s)
	getLogGroup(&c)

}
