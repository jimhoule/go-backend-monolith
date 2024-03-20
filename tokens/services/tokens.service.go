package services

type TokenPayload struct {
	AccountId string
	Email     string
	ExpiresAt int64
}

type TokensService interface {
	GenerateAccessToken(accountId string, email string) (string, error)
	GenerateRefreshToken(accountId string, email string) (string, error)
	Verify(token string) (bool, error)
	Decode(token string) (TokenPayload, error)
}