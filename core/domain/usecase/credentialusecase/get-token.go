package credentialusecase

import (
	"context"
	"time"

	"github.com/sandronister/ibm-cloud-token-go/core/domain"
)

func (u *usecase) GetToken() (*domain.IBMToken, error) {
	cLogin := u.getLoginConfig()
	ctx, cancel := context.WithTimeout(context.Background(), 600*time.Millisecond)
	defer cancel()
	result, err := u.repository.Connect(ctx, cLogin)

	if err != nil {
		return nil, err
	}

	return u.parseToken(result)
}
