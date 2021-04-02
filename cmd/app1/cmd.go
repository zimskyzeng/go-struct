package app1

import (
	"github.com/spf13/cobra"
	"github.com/zimskyzeng/go-struct/config"
)

var (
	AppCmd = &cobra.Command{
		Use:   "",
		Short: "",
		Long:  "",
		PreRun: func(cmd *cobra.Command, args []string) {

		},
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
)

func init() {
	AppCmd.Flags().StringVarP(&config.DefaultConfigPath, "configPath", "c", "",
		"Input ConfingPath")
}
