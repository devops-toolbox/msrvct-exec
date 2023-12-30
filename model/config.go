package model

type CommonConfig struct {
	File      string            `mapstructure:"file" yaml:"file"`
	Path      string            `mapstructure:"path" yaml:"path"`
	Variables map[string]string `mapstructure:"variables" yaml:"variables"`
}
type Config struct {
	Name      string            `mapstructure:"name" yaml:"name"`
	Version   string            `mapstructure:"version" yaml:"version"`
	Label     string            `mapstructure:"label" yaml:"label"`
	Script    Script            `mapstructure:"script" yaml:"script"`
	Variables map[string]string `mapstructure:"variables" yaml:"variables"`
	Packages  []Package         `mapstructure:"packages" yaml:"packages"`
	Common    CommonConfig      `mapstructure:"common" yaml:"common"`
}
type Script struct {
	PreScripts  []string `mapstructure:"pre_scripts" yaml:"pre_scripts"`
	PostScripts []string `mapstructure:"post_scripts" yaml:"post_scripts"`
}

type Package struct {
	Name      string     `yaml:"name"`
	Version   string     `yaml:"version"`
	Type      string     `yaml:"type"`
	Documents []Document `yaml:"documents"`
}
type Document struct {
	Name     string `yaml:"name" mapstructure:"name"`
	Path     string `yaml:"path" mapstructure:"path"`
	Type     string `yaml:"type" mapstructure:"type"`
	SubDir   string `yaml:"sub_dir" mapstructure:"sub_dir"`
	FileMode string `yaml:"file_mode" mapstructure:"file_mode"`
	PathMode string `yaml:"path_mode" mapstructure:"path_mode"`
	Template bool   `yaml:"template" mapstructure:"template" default:"false"`
}
