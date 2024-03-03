package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pandakn/go-sut-course-api/config"
)

type FiberServer struct {
	*fiber.App
	Config config.IConfig
}

func New(cfg config.IConfig) *FiberServer {
	server := &FiberServer{
		App:    fiber.New(),
		Config: cfg,
	}

	return server
}
