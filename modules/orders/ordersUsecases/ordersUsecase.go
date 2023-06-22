package ordersUsecases

import (
	"math"

	"github.com/winai-pgm-itsystem/all-day-shop/modules/entities"
	"github.com/winai-pgm-itsystem/all-day-shop/modules/orders"
	"github.com/winai-pgm-itsystem/all-day-shop/modules/orders/ordersRepositories"
	"github.com/winai-pgm-itsystem/all-day-shop/modules/products/productsRepositories"
)

type IOrdersUsecase interface {
	FindOneOrder(orderId string) (*orders.Order, error)
	FindOrder(req *orders.OrderFilter) *entities.PaginateRes
}

type ordersUsecase struct {
	ordersRepository   ordersRepositories.IOrdersRepository
	productsRepository productsRepositories.IProductsRepository
}

func OrdersUsecase(ordersRepository ordersRepositories.IOrdersRepository, productsRepository productsRepositories.IProductsRepository) IOrdersUsecase {
	return &ordersUsecase{
		ordersRepository:   ordersRepository,
		productsRepository: productsRepository,
	}
}

func (u *ordersUsecase) FindOneOrder(orderId string) (*orders.Order, error) {
	order, err := u.ordersRepository.FindOneOrder(orderId)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (u *ordersUsecase) FindOrder(req *orders.OrderFilter) *entities.PaginateRes {
	orders, count := u.ordersRepository.FindOrder(req)
	return &entities.PaginateRes{
		Data:      orders,
		Page:      req.Page,
		Limit:     req.Limit,
		TotalItem: count,
		TotalPage: int(math.Ceil(float64(count) / float64(req.Limit))),
	}
}
