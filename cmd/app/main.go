package main

import (
	"net/http"

	"github.com/a-shlyapnikov/simple-httprestapi/iternal/database"
	"github.com/a-shlyapnikov/simple-httprestapi/iternal/handlers"
	"github.com/a-shlyapnikov/simple-httprestapi/iternal/messagesService"
	"github.com/gorilla/mux"
)


func main() {
	database.InitDb()
	database.DB.AutoMigrate(&messagesService.Message{})

	repo := messagesService.NewMessageRepository(database.DB)
	service := messagesService.NewService(repo)
	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/get", handler.GetMessagesHandler).Methods("GET")
	router.HandleFunc("/api/post", handler.PostMessageHandler).Methods("POST")
	router.HandleFunc("/api/messages/{id}", handler.UpdateMessageHandler).Methods("PATCH")
	router.HandleFunc("/api/messages/{id}", handler.DeleteMessageHandler).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
