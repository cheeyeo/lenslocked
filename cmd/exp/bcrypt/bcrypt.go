package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func hash(password string) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Printf("error in hash func %v\n", err)
	}
	hash := string(hashedBytes)
	fmt.Println(hash)
}

func compare(password, hash string) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		fmt.Printf("ERR: %s\nPassword is invalid %v\n", err, password)
		return
	}

	fmt.Println("Password is correct")
}

func main() {
	switch os.Args[1] {
	case "hash":
		hash(os.Args[2])
	case "compare":
		compare(os.Args[2], os.Args[3])
	default:
		fmt.Printf("Invalid command: %v\n", os.Args[1])
	}
}
