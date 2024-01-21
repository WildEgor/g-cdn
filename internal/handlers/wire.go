package handlers

import (
	delete_handler "github.com/WildEgor/g-cdn/internal/handlers/delete"
	download_handler "github.com/WildEgor/g-cdn/internal/handlers/download"
	get_files_handler "github.com/WildEgor/g-cdn/internal/handlers/get-files"
	health_check_handler "github.com/WildEgor/g-cdn/internal/handlers/health-check"
	metadata_handler "github.com/WildEgor/g-cdn/internal/handlers/metadata"
	upload_handler "github.com/WildEgor/g-cdn/internal/handlers/upload"
	"github.com/WildEgor/g-cdn/internal/repositories"
	"github.com/google/wire"
)

var HandlersSet = wire.NewSet(
	repositories.RepositoriesSet,
	health_check_handler.NewHealthCheckHandler,
	upload_handler.NewUploadHandler,
	download_handler.NewDownloadHandler,
	delete_handler.NewDeleteHandler,
	metadata_handler.NewMetadataHandler,
	get_files_handler.NewGetFilesHandler,
)
