package credentialusecase

import (
	"encoding/json"

	"github.com/sandronister/ibm-cloud-token-go/core/domain"
)

func (u *usecase) parseToken(body []byte) (*domain.IBMToken, error) {
	token := domain.IBMToken{}
	err := json.Unmarshal(body, &token)
	if err != nil {
		return nil, err
	}
	return &token, nil
}
