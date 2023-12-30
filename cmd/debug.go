/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/devops-toolbox/msrvct-exec/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	DebugCommand := NewDebugCommand()
	rootCmd.AddCommand(DebugCommand)
}
func NewDebugCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "debug",
		Short: "debug",
		Long:  "debug",
		RunE:  DebugRun,
	}
	cmd.Flags().StringSliceP("env", "e", []string{}, "Set environment variables")
	viper.BindPFlag("env", cmd.Flags().Lookup("env"))
	return cmd
}

func DebugRun(cmd *cobra.Command, args []string) (err error) {
	err = internal.DebugRun()
	if err != nil {
		return err
	}
	return nil
}
