package config

type Config interface {
	WordsType() string
	WordsFile() string
	ModelType() string
	ViewType() string
	Ip() string
	Port() string
	ControllerType() string
}
