package metadata_handler

import (
	adapters "github.com/WildEgor/g-cdn/internal/adapters/storage"
	"github.com/WildEgor/g-cdn/internal/dtos"
	core_dtos "github.com/WildEgor/g-core/pkg/core/dtos"
	"github.com/gofiber/fiber/v2"
)

type MetadataHandler struct {
	sa *adapters.StorageAdapter
}

func NewMetadataHandler(
	sa *adapters.StorageAdapter,
) *MetadataHandler {
	return &MetadataHandler{
		sa,
	}
}

func (h *MetadataHandler) Handle(c *fiber.Ctx) error {
	resp := core_dtos.InitResponse()
	resp.FormResponse()

	filename := c.Params("filename")
	if filename == "" {
		resp.SetStatus(c, fiber.StatusBadRequest)
		resp.SetMessage("File name is required")
		resp.JSON(c)
		return nil
	}

	file, err := h.sa.Metadata(filename)
	if err != nil {
		resp.SetStatus(c, fiber.StatusNotFound)
		resp.SetMessage("File not found")
		resp.JSON(c)
		return nil
	}

	dto := &dtos.FileMetadataResponse{
		Filename: filename,
		FileSize: file.Size,
	}
	dto.SetDownloadUrl(c, filename)

	resp.SetStatus(c, fiber.StatusOK)
	resp.SetData([]interface{}{dto})
	resp.JSON(c)
	return nil
}
