package csrf

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func New() fiber.Handler {
	return func(ctx *fiber.Ctx) (err error) {
		verifiedToken := ctx.Locals("verifiedToken").(*jwt.Token)
		if verifiedToken == nil {
			return ctx.SendStatus(fiber.StatusForbidden)
		}

		csrfToken := verifiedToken.Claims.(jwt.MapClaims)["cst"]
		authenticityToken, err := extractFromForm(ctx)
		if err != nil {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}

		if csrfToken == authenticityToken {
			return ctx.Next()
		} else {
			// TODO: expire cookie
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
