package service

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)


func HashingPassword(password string) (string, error)  {
	hashing, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		log.Fatal("Gagal Encrypt Password, Error: ", err.Error())
		return "",err
	}

	return string(hashing),nil

}