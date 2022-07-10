package main

import (
	"database/sql"
	"log"

	"github.com/aryaputra777/rest/config"
	"github.com/aryaputra777/rest/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	if err := config.Connect(); err != nil {
		log.Fatal(err)
	}
	if err = config.Createtableuser(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()
	app.Use(cors.New())

	routes.Listinghandler(app)
	log.Fatal(app.Listen(":4000"))
}
