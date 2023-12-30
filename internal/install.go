package internal

import (
	"github.com/devops-toolbox/msrvct-exec/initialize"
	"github.com/devops-toolbox/msrvct-exec/pkg/install"
)

func InstallRun() (err error) {
	err = initialize.ReadConfig()
	if err != nil {
		return err
	}
	err = install.ExecPreScript()
	if err != nil {
		return err
	}
	err = install.Install()
	if err != nil {
		return err
	}
	err = install.ExecPostScript()
	if err != nil {
		return err
	}
	return nil
}
