package helpers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func CryptPassword(plaintextPassword string) string {
	bytePassword := []byte(plaintextPassword)
	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)

	if (err != nil) {
		log.Println("Error in crypto module: ", err.Error())
	}

	return string(passwordHash)
}
