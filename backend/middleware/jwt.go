package middleware

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func JWTMiddleware() func(http.Handler) http.Handler {

	secretEncoded := os.Getenv("JWT_SECRET")
	secret, err := base64.StdEncoding.DecodeString(secretEncoded)
	if err != nil {
		panic(fmt.Sprintf("Invalid base64 JWT_SECRET: %v", err))
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				http.Error(w, "Missing or invalid token", http.StatusUnauthorized)
				return
			}

			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
			token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method")
				}
				return secret, nil
			})

			if err != nil || !token.Valid {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok || claims["sub"] == nil {
				http.Error(w, "Invalid token claims", http.StatusUnauthorized)
				return
			}

			// Log para verificar se o token chegou e o que contém dentro dele
			log.Printf("Token recebido: %s", tokenStr)
			log.Printf("Claims do token: %+v", claims)

			username := claims["sub"].(string)
			log.Printf("Usuário autenticado: %s", username)

			// Injeta no contexto com a chave "username"
			ctx := context.WithValue(r.Context(), "username", username)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
