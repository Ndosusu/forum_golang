package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// base de données pour stocker les utilisateurs (username et mots de passe hachés)
var users = make(map[string]string)

// Gestionnaire pour la page d'inscription
func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if !goodPasswordLetter(password) {
			http.Error(w, "mdp invalid letter.", http.StatusConflict)
			fmt.Println("erreur mdp")
		}

		if !goodPasswordNumber(password) {
			http.Error(w, "mdp invalid number.", http.StatusConflict)
			fmt.Println("erreur mdp")
		}

		if !goodPasswordExtracharact(password) {
			http.Error(w, "mdp invalid extra.", http.StatusConflict)
			fmt.Println("erreur mdp")
		}

		// Vérifier si l'utilisateur existe déjà
		if _, exists := users[username]; exists {
			http.Error(w, "Cet utilisateur existe déjà.", http.StatusConflict)
			return
		}

		// Hachage du mot de passe
		hashedPassword, err := HashPassword(password)
		if err != nil {
			http.Error(w, "Erreur lors du hachage du mot de passe.", http.StatusInternalServerError)
			return
		}

		// Sauvegarder le nom d'utilisateur et le mot de passe haché
		users[username] = hashedPassword
		fmt.Fprintf(w, "Inscription réussie pour l'utilisateur : %s", username)
	} else {
		http.ServeFile(w, r, "templates/register.html")
	}
}

// Gestionnaire pour la page de connexion
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Récupérer le mot de passe haché pour l'utilisateur
		hashedPassword, exists := users[username]
		if !exists {
			// L'utilisateur n'existe pas
			http.Error(w, "Nom d'utilisateur ou mot de passe incorrect.", http.StatusUnauthorized)
			return
		}

		// Vérifier le mot de passe
		if !CheckPassword(hashedPassword, password) {
			// Mot de passe incorrect
			data := struct {
				Error string
			}{
				Error: "Nom d'utilisateur ou mot de passe incorrect.",
			}
			tmpl := template.Must(template.ParseFiles("templates/login.html"))
			tmpl.Execute(w, data)
			return
		}

		// Connexion réussie
		fmt.Fprintf(w, "Connexion réussie pour l'utilisateur : %s", username)
	} else {
		http.ServeFile(w, r, "templates/login.html")
	}
}
