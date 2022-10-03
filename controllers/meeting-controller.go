package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"client-rest/database"
	"client-rest/models"

	"github.com/gorilla/mux"
)

func GetMeetingById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	key := vars["id"]
	id, err := strconv.Atoi(key)
	if err != nil {
		fmt.Println("error while parsing id value")
	}

	var meeting models.Meeting
	database.Connector.Where("id=?", id).Find(&meeting)

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(meeting)
}

func GetAllMeetings(w http.ResponseWriter, r *http.Request) {
	var meetings []models.Meeting
	database.Connector.Find(&meetings)

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(meetings)
}

func GetClientMeetings(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clientIdStr := vars["clientid"]

	clientId, err := strconv.Atoi(clientIdStr)
	if err != nil {
		fmt.Println("error while parsing id value")
	}

	var meetings []models.Meeting
	database.Connector.Where("clientid=?", clientId).Find(&meetings)

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(meetings)
}

func GetCurrentMeetings(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	currentTimeStr := currentTime.Format("2006-01-02 15:04:05")
	fmt.Println(currentTimeStr)

	var meetings []models.Meeting
	database.Connector.Where("date >= ?", currentTimeStr).Order("date asc").Find(&meetings)

	for i := 0; i < len(meetings); i++ {
		var err error
		date, err := time.Parse("2006-01-02T15:04:05Z", meetings[i].Date)
		if err != nil {
			fmt.Println(err)
		}
		meetings[i].Date = date.Format("2006-01-02 15:04:05")
	}
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(meetings)
}

func CreateMeeting(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var meeting models.Meeting
	json.Unmarshal(requestBody, &meeting)
	database.Connector.Create(meeting)

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(meeting)
}

func DeleteMeetingById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	id, err := strconv.Atoi(key)
	if err != nil {
		fmt.Println("error while parsing id value")
	}

	var meeting models.Meeting
	database.Connector.Where("id=?", id).Delete(&meeting)

	w.WriteHeader(http.StatusNoContent)
}

func UpdateMeetingById(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var meeting models.Meeting
	json.Unmarshal(requestBody, &meeting)
	database.Connector.Save(&meeting)

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(meeting)
}
