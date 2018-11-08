package main

import (
	"encoding/base64"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pwd := `Password123`
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(hash))
	hash2, err2 := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err2 != nil {
		log.Println(err2)
	}
	fmt.Println(string(hash2))

	fmt.Println(base64.StdEncoding.EncodeToString(hash))

	err = bcrypt.CompareHashAndPassword(hash, []byte(`Wrong`))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("OK")
	}

	err = bcrypt.CompareHashAndPassword(hash, []byte(`Password123`))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("OK")
	}
}
