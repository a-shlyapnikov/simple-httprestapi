package main

import (
	"log"

	"github.com/a-shlyapnikov/simple-httprestapi/internal/database"
	"github.com/a-shlyapnikov/simple-httprestapi/internal/handlers"
	"github.com/a-shlyapnikov/simple-httprestapi/internal/messagesService"
	"github.com/a-shlyapnikov/simple-httprestapi/internal/web/messages"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&messagesService.Message{})

	repo := messagesService.NewMessageRepository(database.DB)
	service := messagesService.NewService(repo)

	handler := handlers.NewHandler(service)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := messages.NewStrictHandler(handler, nil)
	messages.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}

	// router := mux.NewRouter()
	// router.HandleFunc("/api/get", handler.GetMessagesHandler).Methods("GET")
	// router.HandleFunc("/api/post", handler.PostMessageHandler).Methods("POST")
	// router.HandleFunc("/api/messages/{id}", handler.UpdateMessageHandler).Methods("PATCH")
	// router.HandleFunc("/api/messages/{id}", handler.DeleteMessageHandler).Methods("DELETE")

	// http.ListenAndServe(":8080", router)
}
