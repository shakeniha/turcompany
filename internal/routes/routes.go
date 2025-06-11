package routes

import (
    "github.com/gorilla/mux"
    "net/http"
    "turcompany/internal/handlers"
    "turcompany/internal/services"
    "turcompany/internal/repositories"
    "database/sql"
)

func NewRouter(db *sql.DB) http.Handler {
    dealRepo := repositories.NewDealRepository(db)
    dealService := services.NewDealService(dealRepo)
    dealHandler := handlers.NewDealHandler(dealService)

    router := mux.NewRouter()
    router.HandleFunc("/deals", dealHandler.Create).Methods("POST")
    router.HandleFunc("/deals/{id}", dealHandler.GetByID).Methods("GET")
    router.HandleFunc("/deals/{id}", dealHandler.Update).Methods("PUT")
    router.HandleFunc("/deals/{id}", dealHandler.Delete).Methods("DELETE")
    return router
}