package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	//you can chose anything you want 
	password := "mohammad"

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)

	fmt.Println(string(hashedPassword))
}
