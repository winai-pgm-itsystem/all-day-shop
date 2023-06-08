package middlewaresUsecases

import (
	"github.com/winai-pgm-itsystem/all-day-shop/modules/middlewares"
	"github.com/winai-pgm-itsystem/all-day-shop/modules/middlewares/middlewaresRepositories"
)

type IMiddlewaresUsecase interface {
	FindAccessToken(userId string, accessToken string) bool
	FindRole() ([]*middlewares.Role, error)
}

type middlewaresUsecase struct {
	middlewaresRepository middlewaresRepositories.IMiddlewaresRepository
}

func MiddlewaresRepository(middlewaresRepository middlewaresRepositories.IMiddlewaresRepository) IMiddlewaresUsecase {
	return &middlewaresUsecase{
		middlewaresRepository: middlewaresRepository,
	}
}

func (u *middlewaresUsecase) FindAccessToken(userId string, accessToken string) bool {
	return u.middlewaresRepository.FindAccessToken(userId, accessToken)
}

func (u *middlewaresUsecase) FindRole() ([]*middlewares.Role, error) {
	roles, err := u.middlewaresRepository.FindRole()
	if err != nil {
		return nil, err
	}
	return roles, nil
}
