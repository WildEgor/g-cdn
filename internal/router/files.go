package router

import (
	delh "github.com/WildEgor/g-cdn/internal/handlers/delete"
	dh "github.com/WildEgor/g-cdn/internal/handlers/download"
	uh "github.com/WildEgor/g-cdn/internal/handlers/upload"
	"github.com/gofiber/fiber/v2"
)

type FilesRouter struct {
	uh   *uh.UploadHandler
	dh   *dh.DownloadHandler
	delh *delh.DeleteHandler
}

func NewFilesRouter(
	uh *uh.UploadHandler,
	dh *dh.DownloadHandler,
	delh *delh.DeleteHandler,
) *FilesRouter {
	return &FilesRouter{
		uh,
		dh,
		delh,
	}
}

func (r *FilesRouter) SetupFilesRouter(app *fiber.App) error {
	v1 := app.Group("/api/v1")

	// TODO: add API-KEY check middleware
	fc := v1.Group("/cdn")

	fc.Post("/upload", r.uh.Handle)
	fc.Get("/download", r.dh.Handle)
	fc.Post("/delete", r.delh.Handle)

	return nil
}
