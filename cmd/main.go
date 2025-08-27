package main

import (
	"crypto-transaction-app/internal/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	routes.SetupRoutes(r)

	// Servir arquivos estáticos da pasta 'public'
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))

	log.Println("Servidor iniciado na porta :8080")
	http.ListenAndServe(":8080", r)
}
