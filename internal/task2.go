package internal

import (
	"errors"
	"fmt"
)

type loginData struct {
	name     string
	password string
}

func Task2(name string, password string) {
	checkValidation, errorMsg := loginValidation(name, password)

	if errorMsg != nil {
		fmt.Println("======Login Failed======")
		fmt.Print("Login Failed : ", errorMsg)
	} else {
		fmt.Println("======Login Success======")
		fmt.Println("Hello !", checkValidation.name)
	}
}

func loginValidation(name string, password string) (loginData, error) {
	userData := loginData{
		name:     "admin",
		password: "admin12345",
	}

	if name != userData.name {
		return userData, errors.New("username salah")
	} else if password != userData.password {
		return userData, errors.New("username salah")
	}

	return userData, nil
}
