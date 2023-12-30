package internal

import (
	"fmt"

	"github.com/devops-toolbox/msrvct-exec/global"
	"github.com/devops-toolbox/msrvct-exec/initialize"
)

func DebugRun() (err error) {
	err = initialize.ReadConfig()
	if err != nil {
		return err
	}
	for k, v := range global.RuntimeVariableMap {
		fmt.Printf("%s: %s\n", k, v)
	}
	// err = install.ExecPreScript()
	// if err != nil {
	// 	return err
	// }
	// err = install.Install()
	// if err != nil {
	// 	return err
	// }
	// err = install.ExecPostScript()
	// if err != nil {
	// 	return err
	// }
	return nil
}
