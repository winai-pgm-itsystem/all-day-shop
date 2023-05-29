package monitorHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/winai-pgm-itsystem/all-day-shop/config"
	"github.com/winai-pgm-itsystem/all-day-shop/modules/entities"
	"github.com/winai-pgm-itsystem/all-day-shop/modules/monitor"
)

type IMonitorHandler interface {
	HealthCheck(c *fiber.Ctx) error
}

type monitorHandlers struct {
	cfg config.IConfig
}

func MonitorHandler(cfg config.IConfig) IMonitorHandler {
	return &monitorHandlers{
		cfg: cfg,
	}
}

func (h *monitorHandlers) HealthCheck(c *fiber.Ctx) error {

	res := &monitor.Monitor{
		Name:    h.cfg.App().Name(),
		Version: h.cfg.App().Version(),
	}
	return entities.NewReponse(c).Success(fiber.StatusOK, res).Res()
}
