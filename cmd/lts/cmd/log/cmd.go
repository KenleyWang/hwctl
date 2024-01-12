package log

import (
	"github.com/spf13/cobra"
	"hwctl/config"
)

var LogCmd = &cobra.Command{
	Use:          "logs",
	Short:        "log",
	SilenceUsage: true,
	Long:         `log`,
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
	getLogs(&c)

}
