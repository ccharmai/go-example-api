package helpers

import (
	"crypto/rsa"
	"encoding/base64"
	"go-example-api/common"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func CryptPassword(plaintextPassword string) string {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), bcrypt.DefaultCost)

	if err != nil {
		log.Println("Error in CryptPassword: ", err.Error())
	}

	return string(passwordHash)
}

func ComparePasswords(hashPassword string, candidatePassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(candidatePassword))

	if err != nil {
		log.Println(err.Error())
	}

	return err == nil
}

func getRSAKeyFromBase64(encodedKey string) *rsa.PrivateKey {
	decodedKey, err := base64.StdEncoding.DecodeString(encodedKey)

	if err != nil {
		log.Println("Base64 decode error: ", err.Error())
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedKey)

	if err != nil {
		log.Println("RSA read error: ", err.Error())
	}

	return key
}

func generateJWTTokenWithPayload(ttl time.Duration, payload interface{}, privateKey string) string {
	decodedPrivateToken := getRSAKeyFromBase64(privateKey)

	now := time.Now().UTC()

	claims := make(jwt.MapClaims)
	claims["exp"] = now.Add(ttl).Unix()
	claims["sub"] = payload

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(decodedPrivateToken)

	if err != nil {
		log.Println("Create JWT error: ", err.Error())
	}

	return token
}

func GenerateAccessToken(payload interface{}) string {
	ttl, _ := time.ParseDuration("10min")
	privateKey := common.Config.AccessTokenPrivateKey
	return generateJWTTokenWithPayload(ttl, payload, privateKey)
}

func GenerateRefreshToken(payload interface{}) string {
	ttl, _ := time.ParseDuration("2d")
	privateKey := common.Config.RefreshTokenPrivateKey
	return generateJWTTokenWithPayload(ttl, payload, privateKey)
}
