package credentialusecase

import "github.com/sandronister/ibm-cloud-token-go/core/domain"

type usecase struct {
	repository domain.CredentialsRepository
}

func New(repository domain.CredentialsRepository) domain.CredentialsUseCase {
	return &usecase{
		repository: repository,
	}
}
