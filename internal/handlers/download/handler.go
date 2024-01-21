package download_handler

import (
	"context"
	adapters "github.com/WildEgor/g-cdn/internal/adapters/storage"
	"github.com/WildEgor/g-cdn/internal/utils"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type DownloadHandler struct {
	sa *adapters.StorageAdapter
}

func NewDownloadHandler(
	sa *adapters.StorageAdapter,
) *DownloadHandler {
	return &DownloadHandler{
		sa,
	}
}

func (hch *DownloadHandler) Handle(c *fiber.Ctx) error {
	ctx := context.Background()

	filename := c.Params("filename")

	file, err := hch.sa.Download(ctx, filename)
	if err != nil {
		return c.SendFile("./public/not_found.png") // TODO
	}

	fBytes := utils.StreamToByte(file)
	if len(fBytes) == 0 {
		return c.SendFile("./public/not_found.png") // TODO
	}

	c.Set("Content-Type", http.DetectContentType(fBytes))

	return c.Send(fBytes)
}
