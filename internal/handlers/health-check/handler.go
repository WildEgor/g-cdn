package health_check_handler

import (
	adapters "github.com/WildEgor/g-cdn/internal/adapters/storage"
	"github.com/WildEgor/g-cdn/internal/config"
	domains "github.com/WildEgor/g-cdn/internal/domain"
	core_dtos "github.com/WildEgor/g-core/pkg/core/dtos"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

type HealthCheckHandler struct {
	st *adapters.StorageAdapter
	ac *config.AppConfig
}

func NewHealthCheckHandler(
	st *adapters.StorageAdapter,
	ac *config.AppConfig,
) *HealthCheckHandler {
	return &HealthCheckHandler{
		st,
		ac,
	}
}

func (hch *HealthCheckHandler) Handle(c *fiber.Ctx) error {
	var health = make(map[string]error)

	resp := core_dtos.InitResponse()

	err := hch.st.Ping()
	if err != nil {
		resp.SetStatus(c, fiber.StatusInternalServerError)
		health["storage"] = err
	}

	for _, h := range health {
		if h != nil {
			log.Error("[HealthCheckHandler] error", h)

			er, _ := health["storage"]
			if er != nil {
				resp.SetError(1, er.Error())
			}
		}
	}

	resp.SetData(&domains.StatusDomain{
		Status:      "ok",
		Version:     hch.ac.Version,
		Environment: hch.ac.GoEnv,
	})
	resp.JSON(c)
	return nil
}
