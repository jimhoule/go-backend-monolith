package services

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtTokensService struct{}

type TokenClaims struct {
	jwt.Claims
	AccountId  string `json:"accountId"`
	Email      string `json:"email"`
	ExpiresAt  int64  `json:"expiresAt"`
}

// TODO: To put in env
var secretKey = []byte("SecretYouShouldHide")

// NOTE: Internal function
func generateToken(accountId string, email string, expiresAt int64) (string, error) {
	// Creates token
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256, 
        jwt.MapClaims{ 
			"accountId": accountId,
			"email": email,
			"expiresAt": expiresAt,
        },
	)

	// Signs token
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (jts *JwtTokensService) GenerateAccessToken(accountId string, email string) (string , error) {
	expiresAt := time.Now().Add(1 * time.Hour).Unix()
	accessToken, err := generateToken(accountId, email, expiresAt)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (jts *JwtTokensService) GenerateRefreshToken(accountId string, email string) (string , error) {
	expiresAt := time.Now().Add(24 * time.Hour).Unix()
	refreshToken, err := generateToken(accountId, email, expiresAt)
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}

func (jts *JwtTokensService) Verify(token string) (bool, error) {
	_, err := jwt.Parse(
		token,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		},
	)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (jts *JwtTokensService) Decode(token string) (TokenPayload , error) {
	tokenPayload := TokenPayload{}

	parsedToken, err := jwt.ParseWithClaims(
		token,
		&TokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		},
	)
	if err != nil {
		return tokenPayload, err
	}

	claims := parsedToken.Claims.(*TokenClaims)
	tokenPayload.AccountId = claims.AccountId
	tokenPayload.Email = claims.Email
	tokenPayload.ExpiresAt = claims.ExpiresAt

	return tokenPayload, nil
}