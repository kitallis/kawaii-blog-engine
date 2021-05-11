package csrf

import (
	"github.com/gofiber/fiber/v2/utils"
)

func Create() string {
	return utils.UUIDv4()
}
