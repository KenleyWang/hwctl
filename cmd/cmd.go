package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"hwctl/cmd/get"
	"hwctl/config"
	"os"
)

func tip() {
	usageStr := `欢迎使用 ` + ` 查看命令`
	usageStr1 := `也可以参考 https://doc.go-admin.dev/guide/ksks 的相关内容`
	fmt.Printf("%s\n", usageStr)
	fmt.Printf("%s\n", usageStr1)
}

func loadConfig(cmd *cobra.Command) (*config.Config, error) {
	var c config.Config
	configPath, err := cmd.Flags().GetString("config")
	if err != nil {
		return &c, err
	}
	err = c.LoadConfig(configPath)
	return &c, nil
}

var rootCmd = &cobra.Command{
	Use:          "hwctl",
	Short:        "hwctl",
	SilenceUsage: true,
	Long:         `hwctl`,
	Version:      `0.5.1`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		//加载配置文件
		c, err := loadConfig(cmd)
		if err != nil {
			return err
		}
		cmd.SetContext(context.WithValue(cmd.Context(), "c", c))
		return err
	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(get.GetCmd)
	rootCmd.PersistentFlags().StringP("config", "c", "./etc/config.yaml", "config file")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(0)
	}
}
