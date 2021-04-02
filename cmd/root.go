package cmd

import "github.com/spf13/cobra"

var (
	RootCmd = &cobra.Command{
		Use: 		"",
		Short: 		"",
		Long: 		"",
	}
)

func init(){
	RootCmd.AddCommand()
}