package middlewares

import (
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/nid90/kawaii-blog-engine/config"
)

type Config struct {
	Name string
	Value string
	Domain string
	Path string
	Expires time.Time
	Secure   bool
	HTTPOnly bool
	SameSite string
}

func defaultConfig() Config {
	return Config{
		Name: "token_",
		Value: "",
		Domain: "",
		Path: "",
		Expires: expirationTime(72),
		Secure: true,
		HTTPOnly: true,
		SameSite: "Strict",
	}
}

func expirationTime(ttl time.Duration) time.Time {
	return time.Now().Add(time.Hour * ttl)
}

func createSignedToken(claims map[string]interface{}) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenClaims := token.Claims.(jwt.MapClaims)
	for key, value := range claims {
		tokenClaims[key] = value
	}
	signedToken, err := token.SignedString([]byte(config.Config("SECRET")))

	return signedToken, err
}

func setCookie(ctx *fiber.Ctx, cfg Config) {
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
}

func SetTokenInCookie(ctx *fiber.Ctx) error {
	// FIXME
	claims := map[string]interface{}{
		"author_id": 123,
		"author_nick": "nick",
		"exp": expirationTime(72).Unix(),
	}
	token, _ := createSignedToken(claims)
	cfg := defaultConfig()
	cfg.Value = token
	setCookie(ctx, cfg)
	return ctx.Next()
}

func UpdateTokenInCookie(ctx *fiber.Ctx) {
	
}