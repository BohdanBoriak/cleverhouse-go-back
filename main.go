package main

import (
	"cleverhouse-go-back/config/container"
	"cleverhouse-go-back/domain"
	"fmt"
	"log"
)

func main() {
	cont := container.New()

	user := domain.User{
		FirstName:  "Bohdan",
		SecondName: "Boriak",
		Phone:      "0123456789",
		Email:      "test@test.com",
	}

	newUser, err := cont.UserRepo.Save(user)
	if err != nil {
		log.Printf("Error: %s", err)
	}

	fmt.Println(newUser)
}
