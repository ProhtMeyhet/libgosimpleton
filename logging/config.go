package logging

import(

)

type DefaultConfig struct {
	Name string
	LogType string
	Level uint8
}

func NewDefaultConfig() *DefaultConfig {
	d := &DefaultConfig{}
	d.LogType = "syslog"
	d.Level = ERRORS
	d.Name = "simpleton"
	return d
}

func (config *DefaultConfig) GetName() string {
	return config.Name
}

func (config *DefaultConfig) GetType() string {
	return config.LogType
}

func (config *DefaultConfig) GetLevel() uint8 {
	return config.Level
}


type DefaultFileConfig struct {
	DefaultConfig
	Path string
}

func NewDefaultFileConfig() *DefaultFileConfig {
	f := &DefaultFileConfig{}
	f.LogType = "file"
	f.Level = ERRORS
	f.Path = "/var/log/simpleton.log"
	f.Name = "simpleton"
	return f
}

func (config *DefaultFileConfig) GetPath() string {
	return config.Path
}
