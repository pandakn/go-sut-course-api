package health

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pandakn/go-sut-course-api/config"
)

type IHealthHandler interface {
	HealthCheck(c *fiber.Ctx) error
}

type healthHandler struct {
	cfg config.IConfig
}

func NewHealthHandler(cfg config.IConfig) IHealthHandler {
	return &healthHandler{
		cfg: cfg,
	}
}

func (h *healthHandler) HealthCheck(c *fiber.Ctx) error {
	res := &Health{
		Name:    h.cfg.App().Name(),
		Message: "everything is fine 🚀",
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
