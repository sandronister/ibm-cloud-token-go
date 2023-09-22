package credentials

import "github.com/gofor-little/env"

type IBMConfig struct {
	Url      string
	GranType string
	ApiKey   string
}

func loadENV() {
	err := env.Load(".env")
	if err != nil {
		panic(err)
	}
}

func LoadConfig() IBMConfig {
	loadENV()
	config := IBMConfig{
		Url:      env.Get("URL", "Fail"),
		GranType: env.Get("GRANT_TYPE", "Fail"),
		ApiKey:   env.Get("API_KEY", "Fail"),
	}
	return config
}
