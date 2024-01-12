package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func tip() {
	usageStr := `欢迎使用 ` + ` 查看命令`
	usageStr1 := `也可以参考 https://doc.go-admin.dev/guide/ksks 的相关内容`
	fmt.Printf("%s\n", usageStr)
	fmt.Printf("%s\n", usageStr1)
}

var rootCmd = &cobra.Command{
	Use:          "hwctl",
	Short:        "hwctl",
	SilenceUsage: true,
	Long:         `hwctl`,
	//PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	//rootCmd.SetHelpTemplate("daw")
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.PersistentFlags().StringP("config", "c", "./etc/config.yaml", "config file")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
