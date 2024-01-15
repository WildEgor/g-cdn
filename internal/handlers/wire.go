package handlers

import (
	health_check_handler "github.com/WildEgor/g-cdn/internal/handlers/health-check"
	"github.com/google/wire"
)

var HandlersSet = wire.NewSet(health_check_handler.NewHealthCheckHandler)
