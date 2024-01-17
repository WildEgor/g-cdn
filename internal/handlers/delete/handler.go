package delete_handler

import (
	core_dtos "github.com/WildEgor/g-core/pkg/core/dtos"
	"github.com/gofiber/fiber/v2"
)

type DeleteHandler struct {
}

func NewDeleteHandler() *DeleteHandler {
	return &DeleteHandler{}
}

func (hch *DeleteHandler) Handle(c *fiber.Ctx) error {
	resp := core_dtos.InitResponse()

	// TODO: delete file from S3 storage and remove meta from mongo db

	resp.FormResponse()
	resp.JSON(c)
	return nil
}
