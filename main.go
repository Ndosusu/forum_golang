package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	port := "8080"

	// Gestionnaire de la route principale "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/index.html")
	})

	// Configuration du serveur HTTP
	server := &http.Server{
		Addr:              ":" + port,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}

	// Lancer le serveur
	fmt.Println("Server starting on http://localhost:" + port)
	if errSrv := server.ListenAndServe(); errSrv != nil {
		log.Fatal(errSrv)
		fmt.Printf("Erreur lors du démarrage du serveur : %v\n", errSrv)
	}
}
