package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"pictionary/internal/database"
	"syscall"

	"github.com/gofiber/fiber/v3"
)

func main() {
	fmt.Println("Hello world")
	var db = database.InitDb()

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error connecting to DB")
	}
	app := fiber.New()

	app.Use(func(c fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound)
	})

	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Fatalf("Error starting server %v", err)
		}
	}()

	shutdown(app, sqlDB)
}

func shutdown(app *fiber.App, sqlDB *sql.DB) {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	fmt.Println("Gracefully shuting down")
	if err := app.Shutdown(); err != nil {
		log.Fatalf("Error shutting down server")
	}

	if err := sqlDB.Close(); err != nil {
		log.Fatalf("Errror closing database connection; %v", err)
	}
}
