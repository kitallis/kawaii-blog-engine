package routes

import (
	"kawaii-blog-engine/handlers"
	"kawaii-blog-engine/middlewares"

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
	authGroup.Get("/new", handlers.SignInView)
	authGroup.Post("/", handlers.SignIn)

	// author
	authorGroup := app.Group("/authors")
	authorGroup.Get("/new", handlers.SignUpView)
	authorGroup.Post("/", handlers.SignUp)

	// post
	postGroup := app.Group("/posts")
	postGroup.Get("/",  middlewares.UpdateTokenInCookie, handlers.FetchPosts)
	postGroup.Post("/", handlers.CreatePost)
	postGroup.Get("/new", handlers.NewPost) // protected
}
