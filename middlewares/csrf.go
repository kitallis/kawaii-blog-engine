package middlewares

// import (
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/gofiber/fiber/v2/middleware/csrf"
// 	"github.com/gofiber/fiber/v2/utils"
// 	"time"
// )

// // Add CSRF protection middleware.
// // Should be done AFTER session middleware.
// var csrfProtection = csrf.New(csrf.Config{
// 	// only to control the switch whether csrf is activated or not
// 	Next: func(c *fiber.Ctx) bool {
// 		return csrfActivated
// 	},
// 	KeyLookup:      "form:_csrf",
// 	CookieName:     "csrf_",
// 	CookieSameSite: "Strict",
// 	Expiration:     1 * time.Hour,
// 	KeyGenerator:   utils.UUID,
// 	ContextKey:     "token",
// })