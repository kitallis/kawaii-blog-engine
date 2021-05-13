package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	jwtCache "kawaii-blog-engine/cache/jwt"
	"kawaii-blog-engine/database"
	"kawaii-blog-engine/models"
	"kawaii-blog-engine/routes"
	"log"
)

func InitMigrations() {
	database.DBConn.AutoMigrate(&models.Post{}, &models.Author{}, &models.Subscriber{})
	fmt.Println("👍🏽 Auto-migrated all models")
}

func main() {
	// init
	database.InitDatabase()
	jwtCache.InitDenyList()
	InitMigrations()
	engine := html.New("./views/", ".html")
	app := fiber.New(fiber.Config{Views: engine})

	// routes
	routes.SetupRoutes(app)

	// start
	log.Fatal(app.Listen(":3000"))
}
