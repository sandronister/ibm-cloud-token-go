package credentialusecase

import (
	"github.com/gofor-little/env"
	"github.com/sandronister/ibm-cloud-token-go/core/domain"
)

func (u *usecase) getLoginConfig() *domain.CloudLogin {
	config := domain.CloudLogin{
		Url:      env.Get("URL", "Fail"),
		GranType: env.Get("GRANT_TYPE", "Fail"),
		ApiKey:   env.Get("API_KEY", "Fail"),
	}
	return &config
}
