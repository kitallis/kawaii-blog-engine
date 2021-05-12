package csrf

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"kawaii-blog-engine/services/cookie"
	"kawaii-blog-engine/services/csrf"
	jwtService "kawaii-blog-engine/services/jwt"
)

func Refresh() fiber.Handler {
	return func(ctx *fiber.Ctx) (err error) {
		csrfToken := csrf.Create()
		verifiedToken := ctx.Locals("verifiedToken").(*jwt.Token)

		claims := verifiedToken.Claims.(jwt.MapClaims)
		claims["cst"] = csrfToken
		newTokenString, _ := jwtService.Create(claims)
		newToken, _ := jwtService.Parse(newTokenString)

		ctx.Cookie(cookie.Create(newTokenString))
		ctx.Locals("CSRFToken", csrfToken)
		ctx.Locals("verifiedToken", newToken)

		return ctx.Next()
	}
}

func Check() fiber.Handler {
	return func(ctx *fiber.Ctx) (err error) {
		verifiedToken := ctx.Locals("verifiedToken").(*jwt.Token)
		csrfToken := verifiedToken.Claims.(jwt.MapClaims)["cst"]
		authenticityToken, err := extractFromForm(ctx)
		if err != nil {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}

		if csrfToken == authenticityToken {
			return ctx.Next()
		} else {
			ctx.Cookie(cookie.Expire())
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
	}
}

const formField = "authenticity_token"

func extractFromForm(ctx *fiber.Ctx) (string, error) {
	token := ctx.FormValue(formField)

	if token == "" {
		return "", errors.New("missing csrf token in form")
	}

	return token, nil
}
