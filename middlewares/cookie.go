package middlewares

import (
	"errors"
	"kawaii-blog-engine/config"
	"kawaii-blog-engine/handlers"
	

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func UpdateTokenInCookie(ctx *fiber.Ctx) error {
	token := ctx.Cookies(config.DefaultCookieConfig().Name)
	if token == "" {
		return errors.New("missing token cookie")
	}
	verifiedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config("SECRET")), nil
	})
	if err != nil {
		ctx.Status(fiber.StatusForbidden)
		return err
	}
	csrfToken := utils.UUID()
	claims := verifiedToken.Claims.(jwt.MapClaims)
	claims["cst"] = csrfToken
	newToken, _ := handlers.CreateSignedToken(claims)

	cfg := config.DefaultCookieConfig()
	cfg.Value = newToken
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
	return ctx.Next()
}