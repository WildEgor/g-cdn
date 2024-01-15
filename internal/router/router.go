package router

import (
	hcHandler "github.com/WildEgor/g-cdn/internal/handlers/health-check"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	hc *hcHandler.HealthCheckHandler
}

func NewRouter(hc *hcHandler.HealthCheckHandler) *Router {
	return &Router{
		hc: hc,
	}
}

func (r *Router) Setup(app *fiber.App) error {
	v1 := app.Group("/api/v1")

	// Server endpoint - sanity check that the server is running
	hcController := v1.Group("/health")
	hcController.Get("/check", r.hc.Handle)

	return nil
}
