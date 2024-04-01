package services

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtTokensService struct{}

type TokenClaims struct {
	jwt.MapClaims
	Subject   string `json:"subject"`
	Email     string `json:"email"`
	IssuedAt  int64  `json:"issuedAt"`
	ExpiresAt int64  `json:"expiresAt"`
}

// TODO: To put in env
var secretKey = []byte("SecretYouShouldHide")

// NOTE: Internal function
func generateToken(accountId string, email string, expiresAt int64) (string, error) {
	// Creates token
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256, 
        TokenClaims{
			Subject: accountId,
			Email: email,
			IssuedAt: time.Now().Unix(),
			ExpiresAt: expiresAt,
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

func (jts *JwtTokensService) Decode(token string) (*TokenPayload , error) {
	tokenPayload := &TokenPayload{}

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
	tokenPayload.Subject = claims.Subject
	tokenPayload.Email = claims.Email
	tokenPayload.IssuedAt =claims.IssuedAt
	tokenPayload.ExpiresAt = claims.ExpiresAt

	return tokenPayload, nil
}