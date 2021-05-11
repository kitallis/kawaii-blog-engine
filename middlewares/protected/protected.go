package protected

import (
	"github.com/gofiber/fiber/v2"
	"kawaii-blog-engine/config"
	jwtService "kawaii-blog-engine/services/jwt"
)

func New() fiber.Handler {
	return func(ctx *fiber.Ctx) (err error) {
		tokenString := ctx.Cookies(config.DefaultCookieConfig().Name)
		verifiedToken, err := jwtService.Parse(tokenString)
		if err != nil {
			return ctx.SendStatus(fiber.StatusForbidden)
		}

		ctx.Locals("verifiedToken", verifiedToken)
		return ctx.Next()
	}
}
