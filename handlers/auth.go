package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/nid90/kawaii-blog-engine/config"
	"github.com/nid90/kawaii-blog-engine/models"
	"golang.org/x/crypto/bcrypt"
	"time"
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

	token := jwt.New(jwt.SigningMethodHS256)
	expiry := time.Now().Add(time.Hour * 72)
	claims := token.Claims.(jwt.MapClaims)
	claims["author_nick"] = author.Nick
	claims["author_id"] = author.ID
	claims["exp"] = expiry.Unix()

	signedToken, err := token.SignedString([]byte(config.Config("SECRET")))

	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "token_",
		Value:    signedToken,
		Domain:   "",
		Path:     "",
		Expires:  expiry,
		Secure:   true,
		HTTPOnly: true,
		SameSite: "Strict",
	})

	return ctx.Redirect("/posts")
}
