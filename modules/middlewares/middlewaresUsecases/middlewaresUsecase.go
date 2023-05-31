package middlewaresUsecases

import (
	"github.com/winai-pgm-itsystem/all-day-shop/modules/middlewares/middlewaresRepositories"
)

type IMiddlewaresUsecase interface{}

type middlewaresUsecase struct {
	middlewaresRepository middlewaresRepositories.IMiddlewaresRepository
}

func MiddlewaresRepository(middlewaresRepository middlewaresRepositories.IMiddlewaresRepository) IMiddlewaresUsecase {
	return &middlewaresUsecase{
		middlewaresRepository: middlewaresRepository,
	}
}
