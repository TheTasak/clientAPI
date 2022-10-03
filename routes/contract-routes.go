package routes

import (
	"client-rest/controllers"

	"github.com/gorilla/mux"
)

func RegisterContractRoutes(router *mux.Router) {
	router.HandleFunc("/contract/", controllers.CreateContract).Methods("POST", "OPTIONS")
	router.HandleFunc("/contract/", controllers.GetAllContracts).Methods("GET", "OPTIONS")
	router.HandleFunc("/contract/{id}", controllers.GetContractById).Methods("GET", "OPTIONS")
	router.HandleFunc("/contract/client/{clientid}", controllers.GetClientContracts).Methods("GET", "OPTIONS")
	router.HandleFunc("/contract/{id}", controllers.UpdateContractById).Methods("PUT", "OPTIONS")
	router.HandleFunc("/contract/{id}", controllers.DeleteContractById).Methods("DELETE", "OPTIONS")
}
