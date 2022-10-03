package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"client-rest/database"
	"client-rest/models"

	"github.com/gorilla/mux"
)

func GetClientById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	key := vars["id"]
	id, err := strconv.Atoi(key)
	if err != nil {
		fmt.Println("error while parsing id value")
	}

	var client models.Client
	database.Connector.Where("id=?", id).Find(&client)

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(client)
}

func GetAllClients(w http.ResponseWriter, r *http.Request) {
	var clients []models.Client
	database.Connector.Find(&clients)

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(clients)
}

func CreateClient(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var client models.Client
	json.Unmarshal(requestBody, &client)
	database.Connector.Create(client)

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(client)
}

func DeleteClientById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	id, err := strconv.Atoi(key)
	if err != nil {
		fmt.Println("error while parsing id value")
	}

	var client models.Client
	database.Connector.Where("id=?", id).Delete(&client)

	w.WriteHeader(http.StatusNoContent)
}

func UpdateClientById(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var client models.Client
	json.Unmarshal(requestBody, &client)
	database.Connector.Save(&client)

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(client)
}
