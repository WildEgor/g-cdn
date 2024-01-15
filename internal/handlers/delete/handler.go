package delete_handler

import (
	domains "github.com/WildEgor/g-cdn/internal/domain"
	"github.com/gofiber/fiber/v2"
)

type DeleteHandler struct {
}

func NewDeleteHandler() *DeleteHandler {
	return &DeleteHandler{}
}

func (hch *DeleteHandler) Handle(c *fiber.Ctx) error {

	// TODO: delete file from S3 storage and remove meta from mongo db

	c.JSON(fiber.Map{
		"isOk": true,
		"data": &domains.StatusDomain{
			Status: "ok",
		},
	})
	return nil
}
