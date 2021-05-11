package handlers

import (
	"kawaii-blog-engine/config"
	"kawaii-blog-engine/models"
	"kawaii-blog-engine/services"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
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

	setTokenInCookie(ctx, author)

	return ctx.Redirect("/posts")
}

func setTokenInCookie(ctx *fiber.Ctx, author *models.Author) error {
	claims := map[string]interface{}{
		"author_id": author.ID,
		"author_nick": author.Nick,
		"exp": config.ExpirationTime(72).Unix(),
	}
	token, err := services.CreateSignedToken(claims)
	if err != nil {
		return err
	}
	cfg := config.DefaultCookieConfig()
	cfg.Value = token
	ctx.Cookie(&fiber.Cookie{
		Name:     cfg.Name,
		Value:    cfg.Value,
		Domain:   cfg.Domain,
		Path:     cfg.Path,
		Expires:  cfg.Expires,
		Secure:   cfg.Secure,
		HTTPOnly: cfg.HTTPOnly,
		SameSite: cfg.SameSite,
	})
	return nil
}
