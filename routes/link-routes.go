package routes

import (
	"client-rest/controllers"

	"github.com/gorilla/mux"
)

func RegisterLinkRoutes(router *mux.Router) {
	router.HandleFunc("/link/", controllers.CreateLink).Methods("POST", "OPTIONS")
	router.HandleFunc("/link/", controllers.GetAllLinks).Methods("GET", "OPTIONS")
	router.HandleFunc("/link/client/{clientid}", controllers.GetClientLinks).Methods("GET", "OPTIONS")
	router.HandleFunc("/link/{id}", controllers.GetLinkById).Methods("GET", "OPTIONS")
	router.HandleFunc("/link/{id}", controllers.UpdateLinkById).Methods("PUT", "OPTIONS")
	router.HandleFunc("/link/{id}", controllers.DeleteLinkById).Methods("DELETE", "OPTIONS")
}
