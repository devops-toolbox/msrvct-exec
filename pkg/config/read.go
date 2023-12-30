package config

import (
	"os"
	"path/filepath"

	"github.com/devops-toolbox/msrvct-exec/global"
	"github.com/devops-toolbox/msrvct-exec/model"
	"github.com/devops-toolbox/msrvct-exec/pkg/utils"
	"github.com/spf13/viper"
)

type Config struct {
	Path string
	File string
}

func NewConfig(path, file string) *Config {
	return &Config{
		Path: path,
		File: file,
	}
}
func (c *Config) ReadConfig() (err error) {
	file := c.File
	if c.File == "" {
		file = global.DefaultConfigFile
	}
	path := file
	if c.Path != "" {
		path = filepath.Join(c.Path, file)
	}
	viper.SetConfigFile(path)
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(global.Config)
	global.RuntimeVariableMap = global.Config.Variables
	if err != nil {
		return err
	}
	return nil
}

func ReadCommonConfig(file, path string) (err error) {

	if file == "" {
		file = global.DefaultConfigFile
	}
	if path != "" {
		path = filepath.Join(path, file)
	} else {
		path = file
	}
	viper.SetConfigFile(path)
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}
	var cc = &model.CommonConfig{}
	err = viper.Unmarshal(cc)
	if err != nil {
		return err
	}
	for k, v := range cc.Variables {
		global.RuntimeVariableMap[k] = v
	}
	return nil
}
func HandleRuntimeVariable() (err error) {
	runtimeVariableList := viper.GetStringSlice("env")
	config := utils.SliceToMap(runtimeVariableList)
	// 覆盖所有指定字段，最高优先级
	for k, v := range config {
		global.RuntimeVariableMap[k] = v
	}
	return nil
}
func HandleDefaultVariable() (err error) {
	config := map[string]string{}
	config["src"], err = os.Getwd()
	if err != nil {
		return err
	}
	config["dst"] = "/usr/local/msrvct"
	config["dst_pkg_dir"] = "pkg"
	config["dst_tmp_dir"] = "tmp"
	config["dst_env_dir"] = "env"
	config["glo_tpl_fix"] = "tpl"
	config["glo_dir_per"] = "0755"
	config["glo_doc_per"] = "0644"
	config["glo_bin_per"] = "0755"
	config["src_scr_dir"] = "scr"
	config["src_pkg_dir"] = "pkg"
	config["src_tpl_dir"] = "tpl"
	// 如果配置中没有指定字段，补齐默认配置，最低优先级
	for k, v := range config {
		if global.RuntimeVariableMap[k] == "" {
			global.RuntimeVariableMap[k] = v
		}
	}
	return nil
}
