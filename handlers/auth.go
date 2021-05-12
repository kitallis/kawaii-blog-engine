package handlers

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"kawaii-blog-engine/config"
	"kawaii-blog-engine/models"
	"kawaii-blog-engine/services/cookie"
	"kawaii-blog-engine/services/csrf"
	jwtService "kawaii-blog-engine/services/jwt"
)

func SignInView(ctx *fiber.Ctx) error {
	return ctx.Render("auth/new", nil, "layout/main")
}

func SignIn(ctx *fiber.Ctx) error {
	type SignInData struct {
		Email    string
		Password string
	}
	
	signInData := new(SignInData)
	if err := ctx.BodyParser(signInData); err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	author, err := models.FindAuthorByEmail(signInData.Email)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	// TODO: gracefully handle password errors
	err = bcrypt.CompareHashAndPassword([]byte(author.Password), []byte(signInData.Password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		ctx.Status(fiber.StatusUnauthorized)
		return SignInView(ctx)
	} else if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	generatedCookie, err := createCookie(author)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	ctx.Cookie(generatedCookie)
	return ctx.Redirect("/posts")
}

func createCookie(author *models.Author) (*fiber.Cookie, error) {
	claims := map[string]interface{}{
		"author_id": author.ID,
		"author_nick": author.Nick,
		"exp": config.ExpirationTime(72).Unix(),
		"cst": csrf.Create(),
		"iat": config.IssueTime(),
	}

	token, err := jwtService.Create(claims)
	if err != nil {
		return nil, err
	}

	return cookie.Create(token), nil
}
