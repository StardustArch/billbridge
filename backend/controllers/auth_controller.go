package controllers

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func AuthInitHandler(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(w, "Missing token", http.StatusUnauthorized)
		return
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

	// Decode secret
	secretEncoded := os.Getenv("JWT_SECRET")
	secret, err := base64.StdEncoding.DecodeString(secretEncoded)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	// Parse JWT
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
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

	username := claims["sub"].(string)

	// Prepara HTML com script para salvar token + claims
	html := fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	<head><title>Autenticando...</title></head>
	<body>
		<script>
			localStorage.setItem("authToken", "%s");
			localStorage.setItem("authUser", JSON.stringify({ username: "%s" }));
			window.location.href = "http://localhost:8081/dashboard";
		</script>
		<p>Redirecionando...</p>
	</body>
	</html>
	`, tokenStr, username)

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}
