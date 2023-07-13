package configs

import (
	"crypto/rand"
	"math/big"
)

const (
	chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	passwordLength = 8
)

func GeneratePassword() string {
	password := make([]byte, passwordLength)
	charSetLength := big.NewInt(int64(len(chars)))

	for i := 0; i < passwordLength; i++ {
		index, _ := rand.Int(rand.Reader, charSetLength)
		password[i] = chars[index.Int64()]
	}

	return string(password)
}