package servers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/winai-pgm-itsystem/all-day-shop/modules/appinfo/appinfoHandlers"
	"github.com/winai-pgm-itsystem/all-day-shop/modules/appinfo/appinfoRepositories"
	"github.com/winai-pgm-itsystem/all-day-shop/modules/appinfo/appinfoUsecases"
	"github.com/winai-pgm-itsystem/all-day-shop/modules/files/filesHandlers"
	"github.com/winai-pgm-itsystem/all-day-shop/modules/files/filesUsecases"
	"github.com/winai-pgm-itsystem/all-day-shop/modules/middlewares/middlewaresHandlers"
	"github.com/winai-pgm-itsystem/all-day-shop/modules/middlewares/middlewaresRepositories"
	"github.com/winai-pgm-itsystem/all-day-shop/modules/middlewares/middlewaresUsecases"
	"github.com/winai-pgm-itsystem/all-day-shop/modules/monitor/monitorHandlers"
	"github.com/winai-pgm-itsystem/all-day-shop/modules/orders/ordersHandlers"
	"github.com/winai-pgm-itsystem/all-day-shop/modules/orders/ordersRepositories"
	"github.com/winai-pgm-itsystem/all-day-shop/modules/orders/ordersUsecases"
	"github.com/winai-pgm-itsystem/all-day-shop/modules/products/productsHandlers"
	"github.com/winai-pgm-itsystem/all-day-shop/modules/products/productsRepositories"
	"github.com/winai-pgm-itsystem/all-day-shop/modules/products/productsUsecases"
	"github.com/winai-pgm-itsystem/all-day-shop/modules/users/usersHandlers"
	"github.com/winai-pgm-itsystem/all-day-shop/modules/users/usersRepositories"
	"github.com/winai-pgm-itsystem/all-day-shop/modules/users/usersUsecases"
)

type IModuleFactory interface {
	MonitorModule()
	UserModule()
	AppinfoModule()
	FilesModule()
	ProductsModule()
	OrdersModule()
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

	router.Post("/signup", m.mid.ApiKeyAuth(), handler.SignUpCustomer)
	router.Post("/signin", m.mid.ApiKeyAuth(), handler.SignIn)
	router.Post("/refresh", m.mid.ApiKeyAuth(), handler.RefreshPassport)
	router.Post("/signout", m.mid.ApiKeyAuth(), handler.SignOut)
	router.Post("/signup-admin", m.mid.JwtAuth(), m.mid.Authorize(2), handler.SignUpAdmin)

	router.Get("/:user_id", m.mid.JwtAuth(), m.mid.ParamsCheck(), handler.GetUserProfile)
	router.Get("/admin/secret", m.mid.JwtAuth(), m.mid.Authorize(2), handler.GenerateAdminToken)
}

func (m *moduleFactory) AppinfoModule() {
	repository := appinfoRepositories.AppinfoRepository(m.s.db)
	usecase := appinfoUsecases.AppinfoUsecase(repository)
	handler := appinfoHandlers.AppinfoHandler(m.s.cfg, usecase)

	router := m.r.Group("/appinfo")

	router.Get("/apikey", m.mid.JwtAuth(), m.mid.Authorize(2), handler.GenerateApiKey)
	router.Get("/categories", m.mid.ApiKeyAuth(), handler.FindCategory)

	router.Post("/categories", m.mid.JwtAuth(), m.mid.Authorize(2), handler.AddCategory)

	router.Delete("/:category_id/categories", m.mid.JwtAuth(), m.mid.Authorize(2), handler.RemoveCategory)

}

func (m *moduleFactory) FilesModule() {

	usecase := filesUsecases.FilesUsecase(m.s.cfg)
	handler := filesHandlers.FilesHandler(m.s.cfg, usecase)

	router := m.r.Group("/files")

	router.Post("/upload", m.mid.JwtAuth(), m.mid.Authorize(2), handler.UploadFiles)
	router.Patch("/delete", m.mid.JwtAuth(), m.mid.Authorize(2), handler.DeleteFile)

}

func (m *moduleFactory) ProductsModule() {

	filesUsecase := filesUsecases.FilesUsecase(m.s.cfg)
	productRepository := productsRepositories.ProductsRepository(m.s.db, m.s.cfg, filesUsecase)
	productsUsecase := productsUsecases.ProductsUsecase(productRepository)
	productsHandler := productsHandlers.ProductsHandler(m.s.cfg, productsUsecase, filesUsecase)

	router := m.r.Group("/products")

	router.Get("/", m.mid.ApiKeyAuth(), productsHandler.FindProduct)
	router.Get("/:product_id", m.mid.ApiKeyAuth(), productsHandler.FindOneProduct)

	router.Post("/", m.mid.JwtAuth(), m.mid.Authorize(2), productsHandler.AddProduct)
	router.Patch("/:product_id", m.mid.JwtAuth(), m.mid.Authorize(2), productsHandler.UpdateProduct)

	router.Delete("/:product_id", m.mid.JwtAuth(), m.mid.Authorize(2), productsHandler.DeleteProduct)

}

func (m *moduleFactory) OrdersModule() {
	filesUsecase := filesUsecases.FilesUsecase(m.s.cfg)
	productRepository := productsRepositories.ProductsRepository(m.s.db, m.s.cfg, filesUsecase)

	ordersRepository := ordersRepositories.OrdersRepository(m.s.db)
	ordersUsecase := ordersUsecases.OrdersUsecase(ordersRepository, productRepository)
	ordersHandler := ordersHandlers.OrdersHandler(m.s.cfg, ordersUsecase)

	router := m.r.Group("/orders")
	router.Get("/:order_id", m.mid.JwtAuth(), ordersHandler.FindOneOrder)

}
