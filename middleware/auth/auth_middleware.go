package authMiddleware

import (
	"encoding/json"
	"insurance-calc-api/context"
	"io"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func ValidateSession(c *fiber.Ctx) error {
	ctx := context.New(c)

	// validate session first
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Create a request object
	req, err := http.NewRequest(http.MethodPost, "http://localhost:3000/api/validate_session", nil)
	if err != nil {
		return ctx.SendInternalServerError(err)
	}

	// Set headers 
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", c.Get("Authorization"))

	// Send the request
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Process the response
	if res.StatusCode != http.StatusOK {
		return ctx.SendErrorMessage(res.StatusCode, "Please make sure you're authorized")
	}

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ctx.SendInternalServerError(err)
	}

	var data = struct {
		UserID uint `json:"user_id"`
	}{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return ctx.SendInternalServerError(err)
	}

	c.Locals("user_id", data.UserID)
	return c.Next()
}
