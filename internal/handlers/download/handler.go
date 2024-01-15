package download_handler

import (
	domains "github.com/WildEgor/g-cdn/internal/domain"
	"github.com/gofiber/fiber/v2"
)

type DownloadHandler struct {
}

func NewDownloadHandler() *DownloadHandler {
	return &DownloadHandler{}
}

func (hch *DownloadHandler) Handle(c *fiber.Ctx) error {

	// TODO: make some logic to check file in db and proxy read from S3 storage
	// TODO: also add functionality to resize and filter images, if no image then just download

	c.JSON(fiber.Map{
		"isOk": true,
		"data": &domains.StatusDomain{
			Status: "ok",
		},
	})
	return nil
}
