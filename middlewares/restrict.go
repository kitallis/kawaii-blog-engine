package middlewares

import "github.com/gofiber/fiber/v2"

func RestrictAccess (ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusForbidden)
}
