package dtos

import "github.com/gofiber/fiber/v2"

type FileResponse struct {
	Filename    string `json:"filename"`
	DownloadUrl string `json:"download_url"`
}

func (fr *FileResponse) SetDownloadUrl(ctx *fiber.Ctx, filename string) {
	fr.DownloadUrl = string(ctx.BaseURL()) + "/api/v1/cdn/download/" + filename
}
