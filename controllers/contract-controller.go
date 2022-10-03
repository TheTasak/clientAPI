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

func GetContractById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	key := vars["id"]
	id, err := strconv.Atoi(key)
	if err != nil {
		fmt.Println("error while parsing id value")
	}

	var contract models.Contract
	database.Connector.Where("id=?", id).Find(&contract)

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(contract)
}

func GetAllContracts(w http.ResponseWriter, r *http.Request) {
	var contacts []models.Contract
	database.Connector.Find(&contacts)

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(contacts)
}

func GetClientContracts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clientIdStr := vars["clientid"]

	clientId, err := strconv.Atoi(clientIdStr)
	if err != nil {
		fmt.Println("error while parsing id value")
	}

	var contracts []models.Contract
	database.Connector.Where("clientid=?", clientId).Find(&contracts)

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(contracts)
}

func CreateContract(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var contract models.Contract
	json.Unmarshal(requestBody, &contract)
	database.Connector.Create(contract)

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(contract)
}

func DeleteContractById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	id, err := strconv.Atoi(key)
	if err != nil {
		fmt.Println("error while parsing id value")
	}

	var contract models.Contract
	database.Connector.Where("id=?", id).Delete(&contract)

	w.WriteHeader(http.StatusNoContent)
}

func UpdateContractById(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var contract models.Contract
	json.Unmarshal(requestBody, &contract)
	database.Connector.Save(&contract)

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(contract)
}
