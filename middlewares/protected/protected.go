package protected

import (
	"github.com/gofiber/fiber/v2"
	"kawaii-blog-engine/config"
	jwtService "kawaii-blog-engine/services/jwt"
)

const rootPath = "/posts"

func Halt() fiber.Handler {
	return func(ctx *fiber.Ctx) (err error) {
		tokenString := ctx.Cookies(config.DefaultCookieConfig().Name)
		newToken, err := jwtService.ParseOrRefresh(tokenString)
		if err != nil {
			return ctx.SendStatus(fiber.StatusForbidden)
		}

		ctx.Locals("verifiedToken", newToken)
		return ctx.Next()
	}
}

func Bypass() fiber.Handler {
	return func(ctx *fiber.Ctx) (err error) {
		tokenString := ctx.Cookies(config.DefaultCookieConfig().Name)
		_, err = jwtService.ParseOrRefresh(tokenString)

		if err != nil {
			return ctx.Next()
		}

		return ctx.Redirect(rootPath)
	}
}
