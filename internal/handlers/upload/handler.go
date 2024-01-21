package upload_handler

import (
	"crypto/md5"
	adapters "github.com/WildEgor/g-cdn/internal/adapters/storage"
	"github.com/WildEgor/g-cdn/internal/repositories"
	"github.com/WildEgor/g-cdn/internal/utils"
	"github.com/WildEgor/g-core/pkg/core/dtos"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"io"
	"mime/multipart"
	"path/filepath"
	"slices"
)

type UploadHandler struct {
	fr *repositories.FileRepository
	sa *adapters.StorageAdapter
}

func NewUploadHandler(
	fr *repositories.FileRepository,
	sa *adapters.StorageAdapter,
) *UploadHandler {
	return &UploadHandler{
		fr,
		sa,
	}
}

func (h *UploadHandler) Handle(c *fiber.Ctx) error {
	ctx := c.Context()
	resp := core_dtos.InitResponse()
	resp.FormResponse()
	log.Debug("Try upload file...")

	var file *multipart.FileHeader
	if form, err := c.MultipartForm(); err == nil {
		file = form.File["files"][0]
	}

	if file == nil {
		resp.SetStatus(c, fiber.StatusBadRequest)
		resp.SetMessage("File not recognized")
		resp.JSON(c)
		return nil
	}

	// TODO: make more readable
	allowedTypes := []string{"image/jpeg", "image/png"}
	contentType := file.Header["Content-Type"][0]

	if slices.Contains(allowedTypes, contentType) != true {
		log.Errorf("File type not allowed: %s", contentType)
		resp.SetStatus(c, fiber.StatusBadRequest)
		resp.SetMessage("File type not allowed")
		resp.JSON(c)
		return nil
	}

	newName := c.FormValue("filename", file.Filename)
	filename := newName + filepath.Ext(file.Filename)
	filteredFilename, err := utils.SanitizeFilename(filename)

	fr, err := file.Open()
	if err != nil {
		resp.SetStatus(c, fiber.StatusBadRequest)
		resp.SetMessage("Failed to open file")
		resp.JSON(c)
		return nil
	}
	defer fr.Close()

	err = h.sa.Upload(ctx, filteredFilename, fr)
	if err != nil {
		resp.SetStatus(c, fiber.StatusInternalServerError)
		resp.SetMessage("Storage upload error")
		resp.JSON(c)
		return nil
	}

	// TODO: make more readable
	fileBuffer := make([]byte, 512)
	for {
		fileBuffer = fileBuffer[:cap(fileBuffer)]
		_, err = fr.Read(fileBuffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			resp.SetStatus(c, fiber.StatusInternalServerError)
			resp.SetMessage("Failed to read file")
			resp.JSON(c)
			break
		}
	}

	checksum := md5.Sum(fileBuffer)

	_, err = h.fr.AddFile(filteredFilename, checksum[:])
	if err != nil {
		log.Errorf(err.Error())
		resp.SetStatus(c, fiber.StatusBadRequest)
		resp.SetMessage("File db save")
		resp.JSON(c)
		return nil
	}

	resp.SetStatus(c, fiber.StatusOK)
	resp.SetMessage("Success")
	resp.JSON(c)

	return nil
}
