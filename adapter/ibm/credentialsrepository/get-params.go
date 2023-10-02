package credentialsrepository

import (
	"net/url"
	"strings"

	"github.com/sandronister/ibm-cloud-token-go/core/domain"
)

func (r *repository) getParams(cLogin domain.CloudLogin) *strings.Reader {
	params := url.Values{}
	params.Add("grant_type", cLogin.GranType)
	params.Add("apikey", cLogin.ApiKey)
	return strings.NewReader(params.Encode())
}
