package error_handler

import (
	"errors"
	core_dtos "github.com/WildEgor/g-core/pkg/core/dtos"
	"github.com/gofiber/fiber/v2"
)

var ErrorHandler = func(c *fiber.Ctx, err error) error {
	resp := core_dtos.InitResponse()

	code := fiber.StatusInternalServerError
	// Retrieve the custom status code if it's a *fiber.TextError
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	headers := make(map[string]string)
	headers[fiber.HeaderContentType] = fiber.MIMEApplicationJSON

	// FIXME: SetHeaders
	c.Set(fiber.HeaderContentType, headers[fiber.HeaderContentType])
	// resp.SetHeaders(c, headers)
	resp.SetStatus(c, code)
	// resp.SetMessage(e.Error())
	resp.JSON(c)

	return nil
}
