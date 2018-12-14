package config

type cliConfig struct {
}

func NewCliConfig() Config {
	return cliConfig{}
}

func (config cliConfig) ServerIp() string {
	return "127.0.0.1"
}
