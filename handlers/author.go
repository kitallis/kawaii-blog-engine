package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nid90/kawaii-blog-engine/database"
	"github.com/nid90/kawaii-blog-engine/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUpView(ctx *fiber.Ctx) error {
	return ctx.Render("author/new", nil, "layout/main")
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func SignUp(ctx *fiber.Ctx) error {
	type SignUpData struct {
		Email     string
		Password  string
		Nick      string
		FirstName string `form:"first_name"`
		LastName  string `form:"last_name"`
	}
	signUpData := new(SignUpData)

	if err := ctx.BodyParser(signUpData); err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return err
	}

	hashedPassword, err := HashPassword(signUpData.Password)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return err
	}
	author := models.Author{
		Email:     signUpData.Email,
		Password:  hashedPassword,
		FirstName: signUpData.FirstName,
		LastName:  signUpData.LastName,
		Nick:      signUpData.Nick,
	}

	database.DBConn.Create(&author)
	return ctx.Redirect("/auth/new")
}