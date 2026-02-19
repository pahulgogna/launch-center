package config

type Config struct {
	AppName string
}

var Env = getConf()

func getConf() Config {
	return Config{
		AppName: "launch-center",
	}
}
