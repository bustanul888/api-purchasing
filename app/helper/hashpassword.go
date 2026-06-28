package helper

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)


func HashPasword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

func ComparePassword(password string, hash string) bool {
	fmt.Println(password)
	fmt.Println(hash)
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
