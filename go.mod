module kawaii-blog-engine

go 1.15

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gofiber/fiber/v2 v2.9.0
	github.com/gofiber/storage/memory v0.0.0-20210507171730-b9b7ed25ca21
	github.com/gofiber/storage/redis v0.0.0-20210507171730-b9b7ed25ca21
	github.com/gofiber/template v1.6.7
	github.com/joho/godotenv v1.3.0
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83
	golang.org/x/sys v0.0.0-20210403161142-5e06dd20ab57 // indirect
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.21.7
)
