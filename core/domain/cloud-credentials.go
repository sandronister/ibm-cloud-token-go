package domain

import "context"

type CloudLogin struct {
	Url      string
	GranType string
	ApiKey   string
}

type IBMToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Expiration   int    `json:"expiration"`
	Scope        string `json:"scope"`
}

type CredentialsUseCase interface {
	GetToken() (*IBMToken, error)
}

type CredentialsRepository interface {
	Connect(ctx context.Context, cLogin *CloudLogin) ([]byte, error)
}

type CredentialsEnviroment interface {
	GetLoginConfig() (*CloudLogin, error)
}
