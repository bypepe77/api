package server

type Config struct {
	AppName string
	Host    string
	Port    string
}

func NewConfig(appName, host string, port string) *Config {
	return &Config{
		AppName: appName,
		Host:    host,
		Port:    port,
	}
}
