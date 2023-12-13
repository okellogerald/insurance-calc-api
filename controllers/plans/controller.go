package plans

import (
	"github.com/gofiber/fiber/v2"
)

func GetPlans(c *fiber.Ctx) error {
	return c.JSON(Plans)
}
