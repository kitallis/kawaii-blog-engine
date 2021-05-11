package cookie

import (
	"github.com/gofiber/fiber/v2"
	"kawaii-blog-engine/config"
)

func Create(value string) *fiber.Cookie {
	cfg := config.DefaultCookieConfig()
	cfg.Value = value

	return &fiber.Cookie{
		Name:     cfg.Name,
		Value:    cfg.Value,
		Domain:   cfg.Domain,
		Path:     cfg.Path,
		Expires:  cfg.Expires,
		Secure:   cfg.Secure,
		HTTPOnly: cfg.HTTPOnly,
		SameSite: cfg.SameSite,
	}
}
