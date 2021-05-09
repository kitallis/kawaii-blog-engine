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

func ExpirationTime(ttl time.Duration) time.Time {
	return time.Now().Add(time.Hour * ttl)
}

func CreateSignedToken(claims map[string]interface{}) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenClaims := token.Claims.(jwt.MapClaims)
	for key, value := range claims {
		tokenClaims[key] = value
	}
	signedToken, err := token.SignedString([]byte(config.Config("SECRET")))

	return signedToken, err
}

func CreateTokenCookie(ctx *fiber.Ctx, token string) {
	ctx.Cookie(&fiber.Cookie{
		Name:     "token_",
		Value:    token,
		Domain:   "",
		Path:     "",
		Expires:  ExpirationTime(72),
		Secure:   true,
		HTTPOnly: true,
		SameSite: "Strict",
	})
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

	claims := map[string]interface{}{
		"author_id": author.ID,
		"author_nick": author.Nick,
		"exp": ExpirationTime(72).Unix(),
	}
	signedToken, err := CreateSignedToken(claims)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	CreateTokenCookie(ctx, signedToken)

	return ctx.Redirect("/posts")
}
