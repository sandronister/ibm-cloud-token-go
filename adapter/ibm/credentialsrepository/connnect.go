package credentialsrepository

import (
	"context"
	"io"
	"net/http"

	"github.com/sandronister/ibm-cloud-token-go/core/domain"
)

func (r *repository) Connect(ctx context.Context, cLogin *domain.CloudLogin) ([]byte, error) {
	req, err := r.getRequest(ctx, *cLogin)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}
