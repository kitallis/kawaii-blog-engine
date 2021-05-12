package routes

import (
	"kawaii-blog-engine/handlers"
	"kawaii-blog-engine/middlewares/cookie"
	"kawaii-blog-engine/middlewares/csrf"
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
	postGroup := app.Group("/posts", protected.New())
	postGroup.Get("/", cookie.Refresh(), handlers.FetchPosts)
	postGroup.Get("/new", cookie.Refresh(), handlers.NewPost)
	postGroup.Post("/", csrf.New(), cookie.Refresh(), handlers.CreatePost)
}
