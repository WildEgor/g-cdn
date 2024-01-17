package upload_handler

import (
	domains "github.com/WildEgor/g-cdn/internal/domain"
	"github.com/WildEgor/g-core/pkg/core/dtos"
	"github.com/gofiber/fiber/v2"
	"slices"
)

type UploadHandler struct {
}

func NewUploadHandler() *UploadHandler {
	return &UploadHandler{}
}

func (hch *UploadHandler) Handle(c *fiber.Ctx) error {

	resp := core_dtos.InitResponse()

	_, err := c.FormFile("files")
	if err != nil {
		return err
	}

	allowedTypes := []string{"image/jpeg", "image/png"}
	if slices.Contains(allowedTypes, fiber.HeaderContentType) != true {
		resp.SetStatus(c, fiber.StatusBadRequest)
		resp.SetMessage("File type not allowed")
		resp.JSON(c)
	}

	// TODO: upload file(s) to S3 storage and save meta to mongo db

	resp.SetStatus(c, fiber.StatusOK)
	resp.SetData(&domains.StatusDomain{
		Status: "ok",
	})
	resp.JSON(c)

	return nil
}
