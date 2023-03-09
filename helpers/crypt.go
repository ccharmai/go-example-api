package helpers

import (
	"crypto/rsa"
	"encoding/base64"
	"errors"
	"fmt"
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

func getRsaPublicKey(encodedKey string) *rsa.PublicKey {
	decodedKey, err := base64.StdEncoding.DecodeString(encodedKey)

	if err != nil {
		log.Println("Base64 decode error: ", err.Error())
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(decodedKey)

	if err != nil {
		log.Println("Parse RSA Private key error: ", err.Error())
	}

	return key
}

func generateJWT(ttl time.Duration, payload interface{}, privateKey string) string {
	rsaPrivate := getRsaPrivateKey(privateKey)

	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().UTC().Add(ttl).Unix()
	claims["sub"] = payload

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(rsaPrivate)

	if err != nil {
		log.Println("Create JWT error: ", err.Error())
	}

	return token
}

func parseJWT(token string, publicKey string) (map[string]interface{}, error) {
	rsaPublic := getRsaPublicKey(publicKey)

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, success := t.Method.(*jwt.SigningMethodRSA); !success {
			return nil, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}
		return rsaPublic, nil
	})

	if err != nil {
		return nil, err
	}

	claims, success := parsedToken.Claims.(jwt.MapClaims)

	if !success || !parsedToken.Valid {
		return nil, errors.New("failed to read token claims")
	}

	payload := claims["sub"].(map[string]interface{})

	return payload, nil
}

type AccessTokenPayload struct {
	ID uint
}

type RefreshTokenPayload struct {
	ID     uint
	IpAddr string
}

func GenerateAccessToken(payload AccessTokenPayload) string {
	ttl, _ := time.ParseDuration("10m")

	privateKey := common.Config.AccessTokenPrivateKey
	return generateJWT(ttl, payload, privateKey)
}

func GenerateRefreshToken(payload RefreshTokenPayload) string {
	ttl, _ := time.ParseDuration("24h")

	privateKey := common.Config.RefreshTokenPrivateKey
	return generateJWT(ttl, payload, privateKey)
}

func getIdFromClaims(payload map[string]interface{}) (uint, bool) {
	parsedId, status := payload["ID"].(float64)

	if !status {
		return 0, false
	}

	return uint(parsedId), true
}

func ParseAccessToken(token string) (AccessTokenPayload, bool) {
	payload := AccessTokenPayload{}

	publicKey := common.Config.AccessTokenPublicKey
	mapClaims, err := parseJWT(token, publicKey)

	if err != nil {
		return payload, false
	}

	id, status := getIdFromClaims(mapClaims)

	if !status {
		return payload, false
	}

	payload.ID = id

	return payload, true
}

func ParseRefreshToken(token string) (RefreshTokenPayload, bool) {
	payload := RefreshTokenPayload{}

	publicKey := common.Config.RefreshTokenPublicKey
	mapClaims, err := parseJWT(token, publicKey)

	if err != nil {
		return payload, false
	}

	id, status := getIdFromClaims(mapClaims)

	if !status {
		return payload, false
	}

	ipAddr, status := mapClaims["IpAddr"].(string)

	if !status {
		return payload, false
	}

	payload.ID = id
	payload.IpAddr = ipAddr

	return payload, true
}
