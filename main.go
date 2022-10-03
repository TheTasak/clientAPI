package main

import (
	"client-rest/database"
	"client-rest/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func initDatabaseConnection() {
	config := database.Config{
		ServerName: "localhost:3306",
		User:       "root",
		Password:   "",
		Database:   "klienci",
	}
	connectionString := config.GetConnectionString()
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
}

func initRoutes(router *mux.Router) {
	routes.RegisterClientRoutes(router)
	routes.RegisterContractRoutes(router)
	routes.RegisterMeetingRoutes(router)
	routes.RegisterLinkRoutes(router)
}

func main() {
	initDatabaseConnection()
	fmt.Println("Starting server on port 8080")

	router := mux.NewRouter().StrictSlash(true)
	initRoutes(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
