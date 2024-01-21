package delete_handler

import (
	adapters "github.com/WildEgor/g-cdn/internal/adapters/storage"
	"github.com/WildEgor/g-cdn/internal/repositories"
	core_dtos "github.com/WildEgor/g-core/pkg/core/dtos"
	"github.com/gofiber/fiber/v2"
)

type DeleteHandler struct {
	fr *repositories.FileRepository
	sa *adapters.StorageAdapter
}

func NewDeleteHandler(
	fr *repositories.FileRepository,
	sa *adapters.StorageAdapter) *DeleteHandler {
	return &DeleteHandler{
		fr,
		sa,
	}
}

// Handle TODO: authorize using X-API-KEY
func (hch *DeleteHandler) Handle(c *fiber.Ctx) error {
	resp := core_dtos.InitResponse()

	filename := c.Params("filename")
	if filename == "" {
		resp.SetStatus(c, fiber.StatusBadRequest)
		resp.SetMessage("File name is required")
		resp.JSON(c)
		return nil
	}

	// TODO
	_ = hch.sa.Delete(filename)
	_, _ = hch.fr.DeleteFile(filename)

	resp.SetStatus(c, fiber.StatusOK)
	resp.SetMessage("Success")
	resp.JSON(c)
	return nil
}
