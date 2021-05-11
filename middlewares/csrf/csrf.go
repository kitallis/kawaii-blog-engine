package csrf

import "github.com/gofiber/fiber/v2"

func New(ctx *fiber.Ctx) error {
	return ctx.Next()
}