package handlers

import (
	delete_handler "github.com/WildEgor/g-cdn/internal/handlers/delete"
	download_handler "github.com/WildEgor/g-cdn/internal/handlers/download"
	health_check_handler "github.com/WildEgor/g-cdn/internal/handlers/health-check"
	upload_handler "github.com/WildEgor/g-cdn/internal/handlers/upload"
	"github.com/google/wire"
)

var HandlersSet = wire.NewSet(
	health_check_handler.NewHealthCheckHandler,
	upload_handler.NewUploadHandler,
	download_handler.NewDownloadHandler,
	delete_handler.NewDeleteHandler,
)
