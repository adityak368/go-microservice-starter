package jwt

import (
	"auth/internal/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// TokenValidityInHours is the validity of the token
const TokenValidityInHours = 72

// GenerateAuthToken Generates Auth Token and sets the claims
func GenerateAuthToken(user *models.User, secret string) (string, error) {

	// Create token
	token := jwt.New(jwt.SigningMethodHS512)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["info"] = user
	claims["sub"] = user.ID
	claims["iss"] = "GoStarter"
	claims["exp"] = time.Now().Add(time.Hour * TokenValidityInHours).Unix()

	// Generate encoded token and send it as response.
	return token.SignedString([]byte(secret))
}
