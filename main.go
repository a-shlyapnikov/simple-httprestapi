package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type requestBody struct {
	Message string `json:"message"`
}

var message string

func main() {
	InitDb()

	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/hello", HelloHandlerPOST).Methods("POST")
	router.HandleFunc("/api/messages", CreateMessage).Methods("POST")
	router.HandleFunc("/api/messages", GetMessage).Methods("GET")
	router.HandleFunc("/api/messages/{id}", UpdateMessage).Methods("PATCH")
	router.HandleFunc("/api/messages/{id}", DeleteMessage).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}

func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	var msg Message
	vars := mux.Vars(r)
	id := vars["id"]

	if result := DB.Delete(&msg, id); result.RowsAffected == 0 {
		http.Error(w, "Message not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func UpdateMessage(w http.ResponseWriter, r *http.Request) {
	var msg Message
	var existingMessage Message
	vars := mux.Vars(r)
	id := vars["id"]

	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if result := DB.First(&existingMessage, id); result.Error != nil || result.RowsAffected == 0 {
		http.Error(w, "Message not found", http.StatusNotFound)
		return
	}

	if err := DB.Model(&existingMessage).Updates(msg).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&msg)
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var msg Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if result := DB.Create(&msg); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&msg)
}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	var messages []Message
	if result := DB.Find(&messages); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)

}

func HelloHandlerPOST(w http.ResponseWriter, r *http.Request) {
	var reqBody requestBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	message = reqBody.Message
	fmt.Fprintf(w, "Hello %s", message)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s", message)
}
