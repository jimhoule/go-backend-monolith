package middlewares

import (
	"app/utils/json"
	"errors"
	"log"
	"net/http"
	"strings"
)

func VerifyAccessToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		authrizationHeader := request.Header.Get("Authorization")

		if authrizationHeader != "" {
			splittedAuthrizationHeader := strings.Split(authrizationHeader, " ")
			authStrategy := splittedAuthrizationHeader[0]
			accessToken := splittedAuthrizationHeader[1]

			// TODO: Make check of strategies const and decode access token
			if authStrategy != "" && accessToken != "" {
				log.Printf("Auth Strategy: %s", authStrategy)
				log.Printf("Access Token: %s", accessToken)

				next.ServeHTTP(writer, request)
				return
			}
		}

		json.WriteHttpError(writer, http.StatusUnauthorized, errors.New("Unauthorized"))
	})
}