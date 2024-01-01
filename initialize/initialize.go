package initialize

import (
	"github.com/devops-toolbox/msrvct-exec/global"
	"github.com/devops-toolbox/msrvct-exec/pkg/config"
)

func ReadConfig() (err error) {
	cc := config.NewConfig(global.DefaultConfigPath, global.DefaultConfigFile)
	err = cc.ReadConfig()
	if err != nil {
		return err
	}
	// log.Println(global.RuntimeVariableMap)
	err = config.HandleRuntimeVariable()
	if err != nil {
		return err
	}
	// log.Println(global.RuntimeVariableMap)
	err = config.ReadCommonConfig()
	if err != nil {
		return err
	}
	err = config.HandleDefaultVariable()
	if err != nil {
		return err
	}
	return nil
}
