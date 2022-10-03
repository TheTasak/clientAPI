package routes

import (
	"client-rest/controllers"

	"github.com/gorilla/mux"
)

func RegisterMeetingRoutes(router *mux.Router) {
	router.HandleFunc("/meeting/", controllers.CreateMeeting).Methods("POST", "OPTIONS")
	router.HandleFunc("/meeting/", controllers.GetAllMeetings).Methods("GET", "OPTIONS")
	router.HandleFunc("/meeting/client/{clientid}", controllers.GetClientMeetings).Methods("GET", "OPTIONS")
	router.HandleFunc("/meeting/current/", controllers.GetCurrentMeetings).Methods("GET", "OPTIONS")
	router.HandleFunc("/meeting/{id}", controllers.GetMeetingById).Methods("GET", "OPTIONS")
	router.HandleFunc("/meeting/{id}", controllers.UpdateMeetingById).Methods("PUT", "OPTIONS")
	router.HandleFunc("/meeting/{id}", controllers.DeleteMeetingById).Methods("DELETE", "OPTIONS")
}
