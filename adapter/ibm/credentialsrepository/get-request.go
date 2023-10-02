package credentialsrepository

import (
	"context"
	"net/http"

	"github.com/sandronister/ibm-cloud-token-go/core/domain"
)

func (r *repository) getRequest(ctx context.Context, cLogin domain.CloudLogin) (*http.Request, error) {
	params := r.getParams(cLogin)
	req, err := http.NewRequestWithContext(ctx, "POST", cLogin.Url, params)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return req, nil
}
