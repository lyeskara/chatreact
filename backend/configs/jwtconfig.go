package configs

import (
	"backend/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWToken(user *models.User) (string, error) {
	JWToken := jwt.New(jwt.SigningMethodHS256)
	claims := JWToken.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 8).Unix()
	JWT, err := JWToken.SignedString([]byte(models.Secret))
	if err != nil {
		return "", err
	}
	return JWT, nil
}


