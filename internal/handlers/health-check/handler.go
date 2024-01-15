package health_check_handler

import (
	adapters "github.com/WildEgor/g-cdn/internal/adapters/storage"
	"github.com/WildEgor/g-cdn/internal/config"
	domains "github.com/WildEgor/g-cdn/internal/domain"
	"github.com/gofiber/fiber/v2"
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
	var health map[string]bool = make(map[string]bool)
	health["storage"] = true

	err := hch.st.Ping()
	if err != nil {
		health["storage"] = false
	}

	for _, h := range health {
		if !h {
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
