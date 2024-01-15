package upload_handler

import (
	domains "github.com/WildEgor/g-cdn/internal/domain"
	"github.com/gofiber/fiber/v2"
)

type UploadHandler struct {
}

func NewUploadHandler() *UploadHandler {
	return &UploadHandler{}
}

func (hch *UploadHandler) Handle(c *fiber.Ctx) error {

	// TODO: upload file(s) to S3 storage and save meta to mongo db

	c.JSON(fiber.Map{
		"isOk": true,
		"data": &domains.StatusDomain{
			Status: "ok",
		},
	})
	return nil
}
