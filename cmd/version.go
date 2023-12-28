/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var (
	_Name_      = ""
	_Version_   = ""
	_GitCommit_ = ""
	_GitBranch_ = ""
	_BuildDate_ = ""
	_BuildTool_ = ""
	OutputType  = "json"
)

func init() {
	versionCommand := NewVersionCommand()
	rootCmd.AddCommand(versionCommand)
}
func NewVersionCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "version",
		Short: "version",
		Long:  "version",
		RunE:  VersionRun,
	}
	cmd.PersistentFlags().StringVarP(&OutputType, "output", "o", OutputType, "Output format. One of: [json yaml]")
	return cmd
}

func VersionRun(cmd *cobra.Command, args []string) (err error) {
	type Program struct {
		Name    string `yaml:"name" json:"name"`
		Version string `yaml:"version" json:"version"`
	}
	type Code struct {
		GitCommit string `yaml:"git_commit" json:"git_commit"`
		GitBranch string `yaml:"git_branch" json:"git_branch"`
	}
	type Build struct {
		BuildDate     string `yaml:"build_date" json:"build_date"`
		BuildTool     string `yaml:"build_tool" json:"build_tool"`
		BuildCompiler string `yaml:"build_compiler" json:"build_compiler"`
	}
	type Version struct {
		Program Program `yaml:"program" json:"program"`
		Code    Code    `yaml:"code" json:"code"`
		Build   Build   `yaml:"build" json:"build"`
		Runtime string  `yaml:"runtime" json:"runtime"`
	}
	versionInfo := Version{
		Program: Program{
			Name:    _Name_,
			Version: _Version_,
		}, Build: Build{
			BuildDate:     _BuildDate_,
			BuildTool:     _BuildTool_,
			BuildCompiler: runtime.Compiler,
		}, Code: Code{
			GitCommit: _GitCommit_,
			GitBranch: _GitBranch_,
		},
		Runtime: runtime.Version(),
	}
	if OutputType == "yaml" {
		versionInfoYaml, err := yaml.Marshal(versionInfo)
		if err != nil {
			return err
		}
		fmt.Println(string(versionInfoYaml))
		return nil
	} else {
		versionInfoJson, err := json.Marshal(versionInfo)
		if err != nil {
			return err
		}
		fmt.Println(string(versionInfoJson))
	}

	return nil
}
