package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = RootCmd()

func RootCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "msrvct-exec",
		Short: "msrvct-exec",
		Long:  "Multiple Software Runtime Version Control Toolbox Executor",
		RunE:  RootRun,
	}
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	return rootCmd
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}

func RootRun(cmd *cobra.Command, args []string) error {
	cmd.Help()
	return nil
}
