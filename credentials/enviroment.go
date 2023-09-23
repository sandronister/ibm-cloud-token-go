package credentials

import "github.com/gofor-little/env"

type Enviroment struct{}

func (e *Enviroment) loadENV() {
	err := env.Load(".env")
	if err != nil {
		panic(err)
	}
}

func (e *Enviroment) GetConfig() *CloudLogin {
	e.loadENV()
	config := CloudLogin{
		Url:      env.Get("URL", "Fail"),
		GranType: env.Get("GRANT_TYPE", "Fail"),
		ApiKey:   env.Get("API_KEY", "Fail"),
	}
	return &config
}
