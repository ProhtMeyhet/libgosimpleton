package logging

import(

)

type Type uint8

const(
	NULL	Type = 0
	BUFFER	Type = 1
	FILE	Type = 2
	SYS	Type = 3
	STDERR	Type = 4
)

type DefaultConfig struct {
	// program name
	Name string
	// type of logger TODO
	LogType Type
	// log level
	Level uint8
	// error handler (name string, e error)
	EHandler func(string, error)

	plain bool
}

func NewDefaultConfig() (config *DefaultConfig) {
	config = &DefaultConfig{}
	config.LogType = SYS
	config.Level = ERRORS
	config.Name = "simpleton"
	config.EHandler = func(name string, e error) {}
	return
}

func (config *DefaultConfig) GetName() string {
	return config.Name
}

func (config *DefaultConfig) GetType() Type {
	return config.LogType
}

func (config *DefaultConfig) GetLevel() uint8 {
	return config.Level
}

func (config *DefaultConfig) SetLevel(to uint8) {
	config.Level = to
}

func (config *DefaultConfig) GetEHandler() func(name string, e error) {
	return config.EHandler
}

func (config *DefaultConfig) SetEHandler(to func(name string, e error)) {
	config.EHandler = to
}

func (config *DefaultConfig) HandleE(name string, e error) {
	config.EHandler(name, e)
}

func (config *DefaultConfig) IsPlain() bool {
	return config.plain
}

func (config *DefaultConfig) TogglePlain() {
	config.plain = !config.plain
}


type DefaultFileConfig struct {
	DefaultConfig
	Path string
}

func NewDefaultFileConfig() (config *DefaultFileConfig) {
	config = &DefaultFileConfig{}
	config.LogType = FILE
	config.Level = ERRORS
	config.Path = "/var/log/simpleton.log"
	config.Name = "simpleton"
	return
}

func (config *DefaultFileConfig) GetPath() string {
	return config.Path
}
