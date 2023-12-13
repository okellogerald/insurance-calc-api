package logs

import (
	"log/slog"
	"os"

	"github.com/gofiber/fiber/v2"
)

var Logger *slog.Logger

func Init() {
	Logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
}

func Info(msg string, args ...any) {
	Logger.Info(msg, args...)
}

func Warn(msg string, args ...any) {
	Logger.Warn(msg, args...)
}

func Error(msg string, args ...any) {
	Logger.Error(msg, args...)
}

func LogRequest(c *fiber.Ctx, args ...any) {
	var l = []any{"method", c.Method(), "path", c.Path()}
	// makes it so that the method and path paramters are before any arguments passed into
	// this function
	args = adjustArgs(l, args...)
	Logger.Info("request-log", args...)
}

func adjustArgs(s []any, str ...any) []any {
	var s2 []any

	s2 = s

	for i := 0; i < len(str); i++ {
		s2 = append(s2, str[i])
	}

	return s2
}
