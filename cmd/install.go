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
	InstallCommand := NewInstallCommand()
	rootCmd.AddCommand(InstallCommand)
}
func NewInstallCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "install",
		Short: "install",
		Long:  "install",
		RunE:  InstallRun,
	}
	cmd.Flags().StringP("config", "c", "", "Set config file")
	viper.BindPFlag("config", cmd.Flags().Lookup("config"))
	cmd.Flags().StringSliceP("env", "e", []string{}, "Set environment variables")
	viper.BindPFlag("env", cmd.Flags().Lookup("env"))
	return cmd
}

func InstallRun(cmd *cobra.Command, args []string) (err error) {
	err = internal.InstallRun()
	if err != nil {
		return err
	}
	return nil
}
