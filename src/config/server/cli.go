package config

import (
	"flag"
)

type cliConfig struct {
	wordsType string
	wordsFile string
	modelType string
	viewType  string
	ip        string
	port      string
}

func NewCliConfig() Config {
	var (
		words string
		file  string
		model string
		view  string
		ip    string
		port  string
	)

	flag.StringVar(&words, "words", "file", "Words Type")
	flag.StringVar(&file, "file", "words/words.txt", "Words file")
	flag.StringVar(&model, "model", "memory", "Model type")
	flag.StringVar(&view, "view", "rest", "View type")
	flag.StringVar(&ip, "ip", "0.0.0.0", "Listening ip")
	flag.StringVar(&port, "port", "8000", "Listening port")
	flag.Parse()

	return cliConfig{
		wordsType: words,
		wordsFile: file,
		modelType: model,
		viewType:  view,
		ip:        ip,
		port:      port}
}

func (config cliConfig) WordsType() string {
	return config.wordsType
}

func (config cliConfig) WordsFile() string {
	return config.wordsFile
}

func (config cliConfig) ModelType() string {
	return config.modelType
}

func (config cliConfig) ViewType() string {
	return config.viewType
}

func (config cliConfig) Ip() string {
	return config.ip
}

func (config cliConfig) Port() string {
	return config.port
}
