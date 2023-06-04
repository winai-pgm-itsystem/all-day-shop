package servers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/winai-pgm-itsystem/all-day-shop/modules/middlewares/middlewaresHandlers"
	"github.com/winai-pgm-itsystem/all-day-shop/modules/middlewares/middlewaresRepositories"
	"github.com/winai-pgm-itsystem/all-day-shop/modules/middlewares/middlewaresUsecases"
	"github.com/winai-pgm-itsystem/all-day-shop/modules/monitor/monitorHandlers"
	"github.com/winai-pgm-itsystem/all-day-shop/modules/users/usersHandlers"
	"github.com/winai-pgm-itsystem/all-day-shop/modules/users/usersRepositories"
	"github.com/winai-pgm-itsystem/all-day-shop/modules/users/usersUsecases"
)

type IModuleFactory interface {
	MonitorModule()
	UserModule()
}

type moduleFactory struct {
	r   fiber.Router
	s   *server
	mid middlewaresHandlers.IMiddlewaresHandler
}

func InitModule(r fiber.Router, s *server, mid middlewaresHandlers.IMiddlewaresHandler) IModuleFactory {
	return &moduleFactory{
		r:   r,
		s:   s,
		mid: mid,
	}
}

func InitMiddlewares(s *server) middlewaresHandlers.IMiddlewaresHandler {
	repository := middlewaresRepositories.MiddlewaresRepository(s.db)
	usecase := middlewaresUsecases.MiddlewaresRepository(repository)
	return middlewaresHandlers.MiddlewaresHandler(s.cfg, usecase)
}

func (m *moduleFactory) MonitorModule() {

	handler := monitorHandlers.MonitorHandler(m.s.cfg)
	m.r.Get("/", handler.HealthCheck)

}

func (m *moduleFactory) UserModule() {
	repository := usersRepositories.UsersRepository(m.s.db)
	usecase := usersUsecases.UsersUsecase(m.s.cfg, repository)
	handler := usersHandlers.UsersHandler(m.s.cfg, usecase)

	router := m.r.Group("/users")

	router.Post("/signup", handler.SignUpCustomer)
	router.Post("/signin", handler.SignIn)
	router.Post("/refresh", handler.RefreshPassport)
}
