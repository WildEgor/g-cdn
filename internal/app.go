package pkg

import (
	"github.com/gofiber/fiber/v2"
	"os"

	"github.com/WildEgor/g-cdn/internal/config"
	"github.com/WildEgor/g-cdn/internal/router"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/google/wire"

	error_handler "github.com/WildEgor/g-cdn/internal/errors"
	log "github.com/sirupsen/logrus"
)

var AppSet = wire.NewSet(
	NewApp,
	config.ConfigsSet,
	router.RouterSet,
)

type Server struct {
	App       *fiber.App
	AppConfig *config.AppConfig
}

func NewApp(
	appConfig *config.AppConfig,
	router *router.Router,
) *Server {
	app := fiber.New(fiber.Config{
		ErrorHandler: error_handler.ErrorHandler,
	})

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin, Content-Type, Accept, Content-Length, Accept-Language, Accept-Encoding, Connection, Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	app.Use(recover.New())

	// Set logging settings
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	if !appConfig.IsProduction() {
		// HINT: some extra setting
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.ErrorLevel)
	}

	router.Setup(app)

	log.Info("Application is running on port...")

	return &Server{
		App:       app,
		AppConfig: appConfig,
	}
}
