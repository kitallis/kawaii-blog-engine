package handlers

import (
	"github.com/gofiber/fiber/v2"
	"kawaii-blog-engine/database"
	"kawaii-blog-engine/models"
)

func CreatePost(ctx *fiber.Ctx) error {
	post := new(models.Post)
	if err := ctx.BodyParser(post); err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return err
	}

	database.DBConn.Create(&post)

	return ctx.Redirect("/posts")
}

type PostsViewData struct {
	Posts []models.Post
}

func FetchPosts(c *fiber.Ctx) error {
	var posts []models.Post
	database.DBConn.Find(&posts)
	return c.Render("post/index", PostsViewData{Posts: posts}, "layout/main")
}

func NewPost(c *fiber.Ctx) error {
	return c.Render("post/new", nil, "layout/main")
}
