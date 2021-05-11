package cookie

import (
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
		if verifiedToken == nil  {
			return ctx.SendStatus(fiber.StatusForbidden)
		}

		claims := verifiedToken.Claims.(jwt.MapClaims)
		claims["cst"] = csrfToken
		newToken, _ := jwtService.Create(claims)
		ctx.Cookie(cookie.Create(newToken))

		return ctx.Next()
	}
}