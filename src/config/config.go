package config

type Config interface {
	ViewType() string
	WordsType() string
	WordsFile() string
	ModelType() string
	RemoteType() string
	Ip() string
	Port() string
}
