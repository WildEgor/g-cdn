package router

import (
	hch "github.com/WildEgor/g-cdn/internal/handlers/health-check"
	"github.com/gofiber/fiber/v2"
)

type HealthRouter struct {
	hch *hch.HealthCheckHandler
}

func NewHealthRouter(hc *hch.HealthCheckHandler) *HealthRouter {
	return &HealthRouter{
		hch: hc,
	}
}

func (r *HealthRouter) SetupHealthRouter(app *fiber.App) error {
	v1 := app.Group("/api/v1")

	hc := v1.Group("/health")
	hc.Get("/check", r.hch.Handle)

	return nil
}
