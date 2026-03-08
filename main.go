package main

import (
	"fmt"
	"net/http"

	"serpgo/server"
)

func main() {
	mux := server.NewMux()

	fmt.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Erreur serveur:", err)
	}
}