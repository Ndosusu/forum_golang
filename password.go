package main

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	// Génération du hash
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
