package download_handler

import (
	core_dtos "github.com/WildEgor/g-core/pkg/core/dtos"
	"github.com/gofiber/fiber/v2"
)

type DownloadHandler struct {
}

func NewDownloadHandler() *DownloadHandler {
	return &DownloadHandler{}
}

func (hch *DownloadHandler) Handle(c *fiber.Ctx) error {
	resp := core_dtos.InitResponse()

	// TODO: make some logic to check file in db and proxy read from S3 storage
	// TODO: also add functionality to resize and filter images, if no image then just download

	resp.FormResponse()
	resp.JSON(c)
	return nil
}
