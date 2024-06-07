package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"dev-docs-cli/pkg/models"
)

// ReadTokenFromFile reads a user token from a specified file
func ReadTokenFromFile(filePath string) (models.User, error) {
	var user models.User

	file, err := os.Open(filePath)
	if err != nil {
		return user, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ": ", 2)
		if len(parts) != 2 {
			continue
		}

		key := parts[0]
		value := parts[1]

		switch key {
		case "AccessToken":
			user.AccessToken = value
		case "Username":
			user.Username = value
		case "Email":
			user.Email = value
		case "UserID":
			user.UserID = value
		}
	}

	if err := scanner.Err(); err != nil {
		return user, err
	}

	return user, nil
}

// IsTokenExpired checks if a JWT token is expired
func IsTokenExpired(tokenString string) (bool, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return true, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if exp, ok := claims["exp"].(float64); ok {
			expirationTime := time.Unix(int64(exp), 0)
			return time.Now().After(expirationTime), nil
		}
	}

	return true, fmt.Errorf("could not parse expiration time")
}

// NewDecodeToken decodes a JWT token and returns its claims as a map
func NewDecodeToken(tokenString string) map[string]interface{} {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		log.Fatalf("Failed to parse token: %v", err)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims
	}
	log.Fatalf("Failed to extract claims from token")
	return nil
}
