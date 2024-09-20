package main

import (
	"golang.org/x/crypto/bcrypt"
)

// func correctPassword(password string) error {
// 	var has bool
// 	for _, l := range password {
// 		if l >= 'a' && l <= 'z' || l >= 'A' && l <= 'Z' {
// 			has = true
// 		}
// 	}

// 	if !has {
// 		return errors.New("mdp invalid letter")
// 	}

// 	has = false
// }

func goodPasswordLetter(password string) bool {
	var letter bool
	for _, l := range password {
		if l >= 'a' && l <= 'z' || l >= 'A' && l <= 'Z' {
			letter = true
		}
	}
	return letter
}

func goodPasswordNumber(password string) bool {
	var number bool
	for _, l := range password {
		if l >= '1' && l <= '9' {
			number = true
		}
	}
	return number
}

func goodPasswordExtracharact(password string) bool {
	var letter bool
	for _, l := range password {
		if l >= 32 && l <= 47 || l >= 58 && l <= 64 || l >= 91 && l <= 96 || l >= 123 && l <= 127 {
			letter = true
		}
	}
	return letter
}

func HashPassword(password string) (string, error) {
	// // Génération du hash
	pass := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost) //"bcrypt.DefaultCost : valeur de complexité du hachage (par défaut 10)"
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
