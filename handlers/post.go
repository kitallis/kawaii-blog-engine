package handlers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	jwtCache "kawaii-blog-engine/cache/jwt"
	"kawaii-blog-engine/database"
	"kawaii-blog-engine/models"
	"kawaii-blog-engine/services/cookie"
	"kawaii-blog-engine/services/csrf"
	jwtService "kawaii-blog-engine/services/jwt"
	"log"
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

func FetchPosts(ctx *fiber.Ctx) error {
	var posts []models.Post
	database.DBConn.Find(&posts)

	log.Println("first time in handler")
	log.Println(jwtCache.DenyList)

	oldTokenString := ctx.Cookies("_kawaii_token")

	log.Println("after cookie read")
	log.Println(jwtCache.DenyList)
	verifiedToken, err := jwtService.Parse(oldTokenString)

	if err != nil {
		fmt.Println(err)
		return ctx.SendStatus(fiber.StatusForbidden)
	}

	csrfToken := csrf.Create()
	claims := verifiedToken.Claims.(jwt.MapClaims)
	claims["cst"] = csrfToken
	newTokenString, _ := jwtService.Create(claims)

	ctx.Cookie(cookie.Create(newTokenString))
	jwtCache.Set(oldTokenString)
	log.Println(jwtCache.DenyList)

	return ctx.Render("post/index", PostsViewData{Posts: posts}, "layout/main")
}

func NewPost(c *fiber.Ctx) error {
	locals := fiber.Map{"authenticity_token": c.Locals("CSRFToken")}
	return c.Render("post/new", locals, "layout/main")
}
