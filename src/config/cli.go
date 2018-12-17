package config

import (
	"flag"
)

type cliConfig struct {
	viewType   string
	wordsType  string
	wordsFile  string
	modelType  string
	remoteType string
	ip         string
	port       string
}

func NewCliConfig() Config {
	var (
		view   string
		words  string
		file   string
		model  string
		remote string
		ip     string
		port   string
	)

	flag.StringVar(&view, "view", "cli", "View Type")
	flag.StringVar(&words, "words", "file", "Words Type")
	flag.StringVar(&file, "file", "words/words.txt", "Words file")
	flag.StringVar(&model, "model", "memory", "Model type")
	flag.StringVar(&remote, "remote", "rest", "Remote connection type")
	flag.StringVar(&ip, "ip", "0.0.0.0", "Listening ip")
	flag.StringVar(&port, "port", "8000", "Listening port")
	flag.Parse()

	return cliConfig{
		viewType:   view,
		wordsType:  words,
		wordsFile:  file,
		modelType:  model,
		remoteType: remote,
		ip:         ip,
		port:       port}
}

func (config cliConfig) ViewType() string {
	return config.viewType
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

func (config cliConfig) RemoteType() string {
	return config.remoteType
}

func (config cliConfig) Ip() string {
	return config.ip
}

func (config cliConfig) Port() string {
	return config.port
}
