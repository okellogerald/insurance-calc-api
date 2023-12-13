package context

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

const (
	VALIDATION_ERROR = "Validation Error"
)

type AuthContext struct {
	FiberCtx *fiber.Ctx
}

func New(c *fiber.Ctx) *AuthContext {
	return &AuthContext{FiberCtx: c}
}

type GlobalResponseMessage struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (c *AuthContext) SendBadRequest(err error) error {
	return c.FiberCtx.Status(fiber.StatusBadRequest).JSON(GlobalResponseMessage{
		Message: err.Error(),
		Success: false,
	})
}

func (c *AuthContext) SendErrorMessage(status int, err string) error {
	return c.FiberCtx.Status(status).JSON(GlobalResponseMessage{
		Message: err,
		Success: false,
	})
}

func (c *AuthContext) SendGlobalResponse(status int, success bool, msg string) error {
	return c.FiberCtx.Status(status).JSON(GlobalResponseMessage{
		Message: msg,
		Success: success,
	})
}

func (c *AuthContext) SendInternalServerError(err error) error {
	return c.FiberCtx.Status(fiber.StatusInternalServerError).JSON(GlobalResponseMessage{
		Message: err.Error(),
		Success: false,
	})
}

func (c *AuthContext) SendUnexpectedInternalError() error {
	return c.FiberCtx.Status(fiber.StatusInternalServerError).JSON(GlobalResponseMessage{
		Message: "Unexpected Internal Error. Please reach out to us immediately",
		Success: false,
	})
}

func (c *AuthContext) SendValidationError(err error) error {
	errors, ok := err.(validator.ValidationErrors)
	if !ok {
		return c.SendBadRequest(err)
	}

	var fields []string
	for _, e := range errors {
		fields = append(fields, e.Field())
	}

	var errorsMap = make(map[string]interface{})
	errorsMap["success"] = false
	errorsMap["message"] = VALIDATION_ERROR
	errorsMap["fields"] = fields

	return c.FiberCtx.Status(fiber.StatusBadRequest).JSON(errorsMap)
}
