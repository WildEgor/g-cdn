package health_check_handler

import (
	adapters "github.com/WildEgor/g-cdn/internal/adapters/storage"
	"github.com/WildEgor/g-cdn/internal/config"
	domains "github.com/WildEgor/g-cdn/internal/domain"
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

	err := hch.st.Ping()
	if err != nil {
		health["storage"] = err
	}

	for _, h := range health {
		if h != nil {
			log.Error("[HealthCheckHandler] error", h)

			c.JSON(fiber.Map{
				"isOk": false,
				"data": &domains.StatusDomain{
					Status:      "fail",
					Version:     hch.ac.Version,
					Environment: hch.ac.GoEnv,
				},
			})
			return nil
		}
	}

	c.JSON(fiber.Map{
		"isOk": true,
		"data": &domains.StatusDomain{
			Status:      "ok",
			Version:     hch.ac.Version,
			Environment: hch.ac.GoEnv,
		},
	})
	return nil
}
