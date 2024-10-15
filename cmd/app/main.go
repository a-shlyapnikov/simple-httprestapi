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
	if err := database.DB.AutoMigrate(&messagesService.Message{}); err != nil{
		log.Fatalf("Auto migration failed: %v",err)
	}

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

}
