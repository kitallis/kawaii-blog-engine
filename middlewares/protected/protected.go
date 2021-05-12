package protected

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"kawaii-blog-engine/config"
	jwtService "kawaii-blog-engine/services/jwt"
	"time"
)

func New() fiber.Handler {
	return func(ctx *fiber.Ctx) (err error) {
		tokenString := ctx.Cookies(config.DefaultCookieConfig().Name)
		verifiedToken, err := jwtService.Parse(tokenString)

		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&(jwt.ValidationErrorExpired) != 0 && verifiedToken != nil {
				claims := verifiedToken.Claims.(jwt.MapClaims)

				refreshUntil, _ := claims["refresh_until"].(json.Number).Int64()
				if refreshUntil > time.Now().Unix() {
					claims["exp"] = config.ExpirationTime(2).Unix()
					newToken, _ := jwtService.Create(claims)
					ctx.Locals("verifiedToken", newToken)
				} else {
					return ctx.SendStatus(fiber.StatusForbidden)
				}
			} else {
				return ctx.SendStatus(fiber.StatusForbidden)
			}
		}

		ctx.Locals("verifiedToken", verifiedToken)
		return ctx.Next()
	}
}
