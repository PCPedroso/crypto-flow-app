package routes

import (
    "net/http"
    "github.com/gorilla/mux"
    "crypto-transaction-app/internal/handlers"
)

func SetupRoutes() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/transaction", handlers.HandleTransaction).Methods("POST")
    router.HandleFunc("/health", handlers.HandleHealthCheck).Methods("GET")

    return router
}