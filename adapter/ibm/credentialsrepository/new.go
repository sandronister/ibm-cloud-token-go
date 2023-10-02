package credentialsrepository

import "github.com/sandronister/ibm-cloud-token-go/core/domain"

type repository struct {
}

func New() domain.CredentialsRepository {
	return &repository{}
}
