package config

import (
	"flag"
)

type Config struct {
	WordsType      string
	WordsFile      string
	ModelType      string
	ViewType       string
	Ip             string
	Port           string
	ControllerType string
}

func NewConfig() (Config, error) {
	var (
		words      string
		file       string
		model      string
		view       string
		ip         string
		port       string
		controller string
	)

	flag.StringVar(&words, "words", "file", "Words Type")
	flag.StringVar(&file, "file", "word/words.txt", "Words file")
	flag.StringVar(&model, "model", "memory", "Model type")
	flag.StringVar(&view, "view", "rest", "View type")
	flag.StringVar(&ip, "ip", "0.0.0.0", "Listening ip")
	flag.StringVar(&port, "port", "8000", "Listening port")
	flag.StringVar(&controller, "controller", "production", "Controller type")
	flag.Parse()

	return Config{
		WordsType:      words,
		WordsFile:      file,
		ModelType:      model,
		ViewType:       view,
		Ip:             ip,
		Port:           port,
		ControllerType: controller}, nil
}
