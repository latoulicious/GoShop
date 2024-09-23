package utils

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateRandomString generates a random string of the specified length
func GenerateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

// GenerateJWTSecretKey generates a new JWT secret key
func GenerateJWTSecretKey() (string, error) {
	secretKey, err := GenerateRandomString(32) // Adjust the length as needed
	if err != nil {
		return "", err
	}
	return secretKey, nil

}
