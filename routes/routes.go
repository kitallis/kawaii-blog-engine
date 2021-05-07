package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/nid90/kawaii-blog-engine/handlers"
)

func SetupRoutes(fiberApp *fiber.App) {
	// public assets
	fiberApp.Static("/", "./assets")

	// middleware
	app := fiberApp.Group("/", logger.New())

	// auth
	authGroup := app.Group("/auth")
	authGroup.Get("/new", handlers.SignInView)
	authGroup.Post("/", handlers.SignIn)
	// authGroup.Delete("/", handlers.SignOut)

	// author
	authorGroup := app.Group("/authors")
	authorGroup.Get("/new", handlers.SignUpView)
	authorGroup.Post("/", handlers.SignUp)

	// post
	postGroup := app.Group("/posts")
	postGroup.Get("/", handlers.FetchPosts)
	postGroup.Post("/", handlers.CreatePost)
	postGroup.Get("/new", handlers.NewPost) // protected
}
