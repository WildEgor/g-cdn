package get_files_handler

import (
	"github.com/WildEgor/g-cdn/internal/config"
	"github.com/WildEgor/g-cdn/internal/dtos"
	"github.com/WildEgor/g-cdn/internal/models"
	"github.com/WildEgor/g-cdn/internal/repositories"
	core_dtos "github.com/WildEgor/g-core/pkg/core/dtos"
	"github.com/gofiber/fiber/v2"
)

type GetFilesHandler struct {
	fr *repositories.FileRepository
	ac *config.AppConfig
}

func NewGetFilesHandler(
	fr *repositories.FileRepository,
	ac *config.AppConfig,
) *GetFilesHandler {
	return &GetFilesHandler{
		fr,
		ac,
	}
}

func (h *GetFilesHandler) Handle(c *fiber.Ctx) error {
	page := c.QueryInt("page")
	if page == 0 || page < 0 {
		page = 1
	}

	limit := c.QueryInt("limit")
	if limit == 0 || limit < 0 || limit > 10 {
		limit = 10
	}

	resp := core_dtos.InitResponse()
	resp.FormResponse()

	pf, err := h.fr.PaginateFiles(&models.PaginationOpts{
		Page:  page,
		Limit: limit,
	})
	if err != nil {
		resp.SetStatus(c, fiber.StatusNotFound)
		resp.SetMessage("Files not found")
		resp.JSON(c)
		return nil
	}

	dto := dtos.NewPaginatedResponse(pf.Total, pf.Data)

	resp.SetStatus(c, fiber.StatusOK)
	resp.SetMessage("Success")
	resp.SetData([]interface{}{dto})
	resp.JSON(c)
	return nil
}
