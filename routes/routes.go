package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"kawaii-blog-engine/handlers"
	"kawaii-blog-engine/middlewares/csrf"
	"kawaii-blog-engine/middlewares/protected"
)

func SetupRoutes(fiberApp *fiber.App) {
	// common middlewares
	app := fiberApp.Group("/", logger.New(logger.Config {
		Format: "[${time}] - ${latency} - ${status} ${method} ${path}\n",
	}))

	// public assets
	app.Static("/", "./assets")

	// auth
	authGroup := app.Group("/auth", protected.Bypass())
	authGroup.Post("/", handlers.SignIn)
	authGroup.Get("/new", handlers.SignInView)

	// author
	authorGroup := app.Group("/authors")
	authorGroup.Post("/", handlers.SignUp)
	authorGroup.Get("/new", handlers.SignUpView)

	// post
	postGroup := app.Group("/posts")
	postGroup.Post("/", csrf.Check(), csrf.Refresh(), handlers.CreatePost)
	postGroup.Get("/", handlers.FetchPosts)
	postGroup.Get("/new", csrf.Refresh(), handlers.NewPost)
}