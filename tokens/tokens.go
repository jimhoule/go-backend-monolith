package tokens

import "app/tokens/services"

func GetService() services.TokensService {
	return &services.JwtTokensService{}
}