package model

import (
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/i-jonathan/pharmacy-api/config"
	appError "github.com/i-jonathan/pharmacy-api/error"
	"golang.org/x/crypto/argon2"
	"strings"
	"time"
)

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a *Account) CreateToken() (string, error) {
	hash, err := ToHashID(a.ID)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"hash":           hash,
		"StandardClaims": jwt.StandardClaims{ExpiresAt: time.Now().Add(168 * time.Hour).Unix()},
	})
	config2 := config.GetConfig()
	hmacSecret := []byte(config2.HMAC)
	tokenString, err := token.SignedString(hmacSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (a *Auth) ComparePassword(hash string) (bool, error) {
	parts := strings.Split(hash, "$")
	passConfig := &passwordConfig{}

	_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &passConfig.memory, &passConfig.time, &passConfig.threads)
	if err != nil {
		return false, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}

	decodedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}

	passConfig.keyLen = uint32(len(decodedHash))

	comparisonHash := argon2.IDKey([]byte(a.Password), salt, passConfig.time, passConfig.memory, passConfig.threads, passConfig.keyLen)

	return subtle.ConstantTimeCompare(decodedHash, comparisonHash) == 1, nil
}

func ParseToken(tokenString string) (map[string]interface{}, error) {
	// TODO check if token is blacklisted
	config2 := config.GetConfig()
	hmacSecret := []byte(config2.HMAC)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, appError.Unauthorized
		}
		return hmacSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
