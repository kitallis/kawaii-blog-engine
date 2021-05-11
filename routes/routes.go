package routes

import (
	"kawaii-blog-engine/handlers"
	"kawaii-blog-engine/middlewares/cookie"
	"kawaii-blog-engine/middlewares/protected"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(fiberApp *fiber.App) {
	// public assets
	fiberApp.Static("/", "./assets")

	// common middlewares
	app := fiberApp.Group("/", logger.New())

	// auth
	authGroup := app.Group("/auth")
	authGroup.Post("/", handlers.SignIn)
	authGroup.Get("/new", handlers.SignInView)

	// author
	authorGroup := app.Group("/authors")
	authorGroup.Post("/", handlers.SignUp)
	authorGroup.Get("/new", handlers.SignUpView)

	// post
	postGroup := app.Group("/posts", protected.New(), cookie.Refresh())
	postGroup.Get("/", handlers.FetchPosts)
	postGroup.Get("/new", handlers.NewPost)
	postGroup.Post("/", handlers.CreatePost)
}
