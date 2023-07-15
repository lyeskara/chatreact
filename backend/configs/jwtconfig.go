package configs

import (
	"github.com/golang-jwt/jwt/v5"
	"backend/models"
	"time"
)

func GenerateJWToken(user *models.User) (string, error) {
	JWToken := jwt.New(jwt.SigningMethodHS256)
	claims := JWToken.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 8).Unix()
	JWT, err := JWToken.SignedString([]byte("KDks9v2f9EK1B1RZ"))
	if err != nil {
		return "", err
	}
	return JWT, nil
}
