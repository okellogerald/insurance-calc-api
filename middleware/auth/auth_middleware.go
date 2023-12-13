package authMiddleware

import (
	"encoding/json"
	"insurance-calc-api/context"

	"github.com/gofiber/fiber/v2"
)

func ValidateSession(c *fiber.Ctx) error {
	ctx := context.New(c)
	// validate session first
	agent := fiber.Post("http://localhost:3000/api/validate_session")

	token := c.Get("Authorization")
	agent.Request().Header.Add("Authorization", token)
	statusCode, body, errs := agent.Bytes()
	if statusCode != 200 {
		return ctx.SendErrorMessage(statusCode, "Please make sure you're authorized")
	}
	if len(errs) > 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errs": errs,
		})
	}

	var data = struct {
		UserID uint `json:"user_id"`
	}{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		return ctx.SendInternalServerError(err)
	}

	c.Locals("user_id", data.UserID)
	return c.Next()
}
