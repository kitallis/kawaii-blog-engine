package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nid90/kawaii-blog-engine/models"
	"golang.org/x/crypto/bcrypt"
)

func SignInView(ctx *fiber.Ctx) error {
	return ctx.Render("auth/new", nil, "layout/main")
}

func SignIn(ctx *fiber.Ctx) error {
	type SignInData struct {
		Email    string
		Password string
	}
	
	signIndata := new(SignInData)
	if err := ctx.BodyParser(signIndata); err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return err
	}

	author, err := models.FindAuthorByEmail(signIndata.Email)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(author.Password), []byte(signIndata.Password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		// TODO: gracefully handle password errors
		ctx.Status(fiber.StatusUnauthorized)
		return SignInView(ctx)
	} else if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return err
	}

	return ctx.Redirect("/posts")
}
