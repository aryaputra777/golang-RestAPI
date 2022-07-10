package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// GenerateNewAccessToken func for generate a new Access token.
func GenerateNewAccessToken(user_id uint32) (string, error) {
	// Set secret key from .env file.
	secret := os.Getenv("JWT_SECRET_KEY")

	// Set expires minutes count for secret key from .env file.
	//minutesCount, _ := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))

	// Create a new claims.
	claims := jwt.MapClaims{}

	// Set public claims:
	claims["exp"] = time.Now().Add(time.Hour * 3).Unix() //Token expires after 1 hour
	claims["authorized"] = true
	claims["user_id"] = user_id
	// Create a new JWT access token with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate token.
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		// Return error, it JWT token generation failed.
		return "", err
	}

	return t, nil
}
