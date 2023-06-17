package productsUsecases

import (
	"github.com/winai-pgm-itsystem/all-day-shop/modules/products"
	"github.com/winai-pgm-itsystem/all-day-shop/modules/products/productsRepositories"
)

type IProductsUsecase interface {
	FindOneProduct(productId string) (*products.Product, error)
}

type productsUsecase struct {
	productsRepositories productsRepositories.IProductsRepository
}

func ProductsUsecase(productsRepositories productsRepositories.IProductsRepository) IProductsUsecase {
	return &productsUsecase{
		productsRepositories: productsRepositories,
	}
}

func (u *productsUsecase) FindOneProduct(productId string) (*products.Product, error) {
	product, err := u.productsRepositories.FindOneProduct(productId)
	if err != nil {
		return nil, err
	}
	return product, nil
}
