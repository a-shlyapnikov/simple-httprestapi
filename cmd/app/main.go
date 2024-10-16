package main

import (
	"log"

	"github.com/a-shlyapnikov/simple-httprestapi/internal/database"
	"github.com/a-shlyapnikov/simple-httprestapi/internal/handlers"
	"github.com/a-shlyapnikov/simple-httprestapi/internal/messagesService"
	"github.com/a-shlyapnikov/simple-httprestapi/internal/userService"
	"github.com/a-shlyapnikov/simple-httprestapi/internal/web/messages"
	"github.com/a-shlyapnikov/simple-httprestapi/internal/web/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database.InitDB()
	if err := database.DB.AutoMigrate(&messagesService.Message{}, &userService.User{}); err != nil {
		log.Fatalf("Auto migration failed: %v", err)
	}

	messagesRepo := messagesService.NewMessageRepository(database.DB)
	messagesService := messagesService.NewService(messagesRepo)
	messagesHandler := handlers.NewMessageHandler(messagesService)

	usersRepo := userService.NewUserRepository(database.DB)
	usersService := userService.NewService(usersRepo)
	usersHandler := handlers.NewUserHandler(usersService)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	usersStrictHandler := users.NewStrictHandler(usersHandler, nil)
	users.RegisterHandlers(e, usersStrictHandler)

	messagesStrictHandler := messages.NewStrictHandler(messagesHandler, nil)
	messages.RegisterHandlers(e, messagesStrictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}

}
