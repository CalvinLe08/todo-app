package utils

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(ttl time.Duration, payload interface{}, privateKey string) (token string, err error) {
	decodedKey, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return  "", fmt.Errorf("Could not decode key, %w", err)
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedKey)
	if err != nil {
		return "", fmt.Errorf("create: parse key: %w", err)
	}

	now := time.Now().UTC()

	claims := make(jwt.MapClaims)
	claims["sub"] = payload
	claims["exp"] = now.Add(ttl).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	token, err = jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)

	if err != nil {
		return "", fmt.Errorf("create: sign token: %w", err)
	}

	return token, nil
}
