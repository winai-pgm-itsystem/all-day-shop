package productsHandlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/winai-pgm-itsystem/all-day-shop/config"
	"github.com/winai-pgm-itsystem/all-day-shop/modules/entities"
	"github.com/winai-pgm-itsystem/all-day-shop/modules/files/filesUsecases"
	"github.com/winai-pgm-itsystem/all-day-shop/modules/products/productsUsecases"
)

type productsHandelersErrCode string

const (
	findOneProductErr productsHandelersErrCode = "products-001"
)

type IProductsHandler interface {
	FindOneProduct(c *fiber.Ctx) error
}

type productsHandler struct {
	cfg              config.IConfig
	productsUsecases productsUsecases.IProductsUsecase
	filesUsecases    filesUsecases.IFilesUsecase
}

func ProductsHandler(cfg config.IConfig, productsUsecases productsUsecases.IProductsUsecase, filesUsecases filesUsecases.IFilesUsecase) IProductsHandler {
	return &productsHandler{
		cfg:              cfg,
		productsUsecases: productsUsecases,
		filesUsecases:    filesUsecases,
	}
}

func (h *productsHandler) FindOneProduct(c *fiber.Ctx) error {

	productId := strings.Trim(c.Params("product_id"), " ")

	product, err := h.productsUsecases.FindOneProduct(productId)
	if err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrInternalServerError.Code,
			string(findOneProductErr),
			err.Error(),
		).Res()
	}

	return entities.NewResponse(c).Success(fiber.StatusOK, product).Res()
}
