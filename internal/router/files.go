package router

import (
	delh "github.com/WildEgor/g-cdn/internal/handlers/delete"
	dh "github.com/WildEgor/g-cdn/internal/handlers/download"
	gf "github.com/WildEgor/g-cdn/internal/handlers/get-files"
	mh "github.com/WildEgor/g-cdn/internal/handlers/metadata"
	uh "github.com/WildEgor/g-cdn/internal/handlers/upload"
	"github.com/gofiber/fiber/v2"
)

type FilesRouter struct {
	uh   *uh.UploadHandler
	dh   *dh.DownloadHandler
	delh *delh.DeleteHandler
	gf   *gf.GetFilesHandler
	mh   *mh.MetadataHandler
}

func NewFilesRouter(
	uh *uh.UploadHandler,
	dh *dh.DownloadHandler,
	delh *delh.DeleteHandler,
	gf *gf.GetFilesHandler,
	mh *mh.MetadataHandler,
) *FilesRouter {
	return &FilesRouter{
		uh,
		dh,
		delh,
		gf,
		mh,
	}
}

func (r *FilesRouter) SetupFilesRouter(app *fiber.App) error {
	v1 := app.Group("/api/v1")

	// TODO: add API-KEY check middleware
	fc := v1.Group("/cdn")

	fc.Post("/upload", r.uh.Handle)
	fc.Get("/download/:filename", r.dh.Handle)
	fc.Delete("/delete/:filename", r.delh.Handle)
	fc.Get("/metadata/:filename", r.mh.Handle)
	fc.Get("/files", r.gf.Handle)

	return nil
}
