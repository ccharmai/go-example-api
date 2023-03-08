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

func HashPassword(plaintextPassword string) string {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), bcrypt.DefaultCost)

	if err != nil {
		log.Println("Error in CryptPassword: ", err.Error())
	}

	return string(passwordHash)
}

func CompareHashAndPlaintextPassword(hashPassword string, candidatePassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(candidatePassword))

	return err == nil
}

func getRsaPrivateKey(encodedKey string) *rsa.PrivateKey {
	decodedKey, err := base64.StdEncoding.DecodeString(encodedKey)

	if err != nil {
		log.Println("Base64 decode error: ", err.Error())
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedKey)

	if err != nil {
		log.Println("Parse RSA Private key error: ", err.Error())
	}

	return key
}

func generateJWT(ttl time.Duration, payload interface{}, privateKey string) string {
	decodedPrivateKey := getRsaPrivateKey(privateKey)

	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().UTC().Add(ttl).Unix()
	claims["sub"] = payload

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(decodedPrivateKey)

	if err != nil {
		log.Println("Create JWT error: ", err.Error())
	}

	return token
}

type AccessTokenPayload struct {
	Email string
}

type RefreshTokenPayload struct {
	Email  string
	IpAddr string
}

func GenerateAccessToken(payload AccessTokenPayload) string {
	ttl, _ := time.ParseDuration("10min")
	privateKey := common.Config.AccessTokenPrivateKey
	return generateJWT(ttl, payload, privateKey)
}

func GenerateRefreshToken(payload RefreshTokenPayload) string {
	ttl, _ := time.ParseDuration("2d")
	privateKey := common.Config.RefreshTokenPrivateKey
	return generateJWT(ttl, payload, privateKey)
}
