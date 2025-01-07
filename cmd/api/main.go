package main

import (
	"fmt"
	"github.com/MrScotch679/book-tickets-backend/config"
	"github.com/MrScotch679/book-tickets-backend/db"
	"github.com/MrScotch679/book-tickets-backend/handlers"
	"github.com/MrScotch679/book-tickets-backend/repositories"
	"github.com/gofiber/fiber/v2"
)

func main() {
	envConfig := config.NewEnvConfig()
	initDb := db.InitDb(envConfig, db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName:      "BookTickets",
		ServerHeader: "Fiber",
	})

	eventRepository := repositories.NewEventRepository(initDb)

	server := app.Group("/api/v1")

	handlers.NewEventHandler(server.Group("/event"), eventRepository)

	app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}
