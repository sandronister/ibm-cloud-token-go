package di

import (
	"github.com/sandronister/ibm-cloud-token-go/adapter/ibm/credentialsrepository"
	"github.com/sandronister/ibm-cloud-token-go/core/domain"
	"github.com/sandronister/ibm-cloud-token-go/core/domain/usecase/credentialusecase"
)

func ConfigIBMConect() domain.CredentialsUseCase {
	repository := credentialsrepository.New()
	return credentialusecase.New(repository)
}
