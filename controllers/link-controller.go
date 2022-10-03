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

func GetLinkById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	key := vars["id"]
	id, err := strconv.Atoi(key)
	if err != nil {
		fmt.Println("error while parsing id value")
	}

	var link models.Link
	database.Connector.Where("id=?", id).Find(&link)

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(link)
}

func GetAllLinks(w http.ResponseWriter, r *http.Request) {
	var links []models.Link
	database.Connector.Find(&links)

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(links)
}

func GetClientLinks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clientIdStr := vars["clientid"]

	clientId, err := strconv.Atoi(clientIdStr)
	if err != nil {
		fmt.Println("error while parsing id value")
	}

	var links []models.Link
	database.Connector.Where("idsource=?", clientId).Find(&links)

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(links)
}

func CreateLink(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var link models.Link
	json.Unmarshal(requestBody, &link)
	database.Connector.Create(link)

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(link)
}

func DeleteLinkById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	id, err := strconv.Atoi(key)
	if err != nil {
		fmt.Println("error while parsing id value")
	}

	var link models.Link
	database.Connector.Where("id=?", id).Delete(&link)

	w.WriteHeader(http.StatusNoContent)
}

func UpdateLinkById(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var link models.Link
	json.Unmarshal(requestBody, &link)
	database.Connector.Save(&link)

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(link)
}
