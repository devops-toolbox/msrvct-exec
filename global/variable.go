package global

import "github.com/devops-toolbox/msrvct-exec/model"

var DefaultConfigPath = ""
var DefaultConfigFile = "config.yaml"
var Config = &model.Config{}

var DefaultCommonConfigPath = "/etc/msrvct"
var DefaultCommonConfigFile = "msrvct.yaml"
var RuntimeVariableMap = map[string]string{}
