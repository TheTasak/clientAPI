package routes

import (
	"client-rest/controllers"

	"github.com/gorilla/mux"
)

func RegisterClientRoutes(router *mux.Router) {
	router.HandleFunc("/client/", controllers.CreateClient).Methods("POST", "OPTIONS")
	router.HandleFunc("/client/", controllers.GetAllClients).Methods("GET", "OPTIONS")
	router.HandleFunc("/client/{id}", controllers.GetClientById).Methods("GET", "OPTIONS")
	router.HandleFunc("/client/{id}", controllers.UpdateClientById).Methods("PUT", "OPTIONS")
	router.HandleFunc("/client/{id}", controllers.DeleteClientById).Methods("DELETE", "OPTIONS")
}
