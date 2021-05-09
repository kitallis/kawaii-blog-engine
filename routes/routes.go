package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/nid90/kawaii-blog-engine/handlers"
	"github.com/nid90/kawaii-blog-engine/middlewares"
)

func SetupRoutes(fiberApp *fiber.App) {
	// public assets
	fiberApp.Static("/", "./assets")

	// middleware
	app := fiberApp.Group("/", logger.New())
	// app = app.Group("/posts", middlewares.RestrictAccess)

	// auth
	authGroup := app.Group("/auth")
	authGroup.Get("/new", handlers.SignInView)
	authGroup.Post("/", middlewares.SetTokenInCookie, handlers.SignIn)
	// authGroup.Delete("/", handlers.SignOut)x

	// author
	authorGroup := app.Group("/authors")
	authorGroup.Get("/new", handlers.SignUpView)
	authorGroup.Post("/", handlers.SignUp)

	// post
	postGroup := app.Group("/posts", middlewares.RestrictAccess)
	postGroup.Get("/", handlers.FetchPosts)
	postGroup.Post("/", handlers.CreatePost)
	postGroup.Get("/new", handlers.NewPost) // protected
}
