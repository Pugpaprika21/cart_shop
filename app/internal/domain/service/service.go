package service

import (
	"miniservice/app/internal/adapter/repository"
	"miniservice/app/internal/domain/service/products"
	"miniservice/app/internal/domain/service/user"
)

type Service struct {
	User     user.IUserService
	Products products.IProductsService
}

func New(repository *repository.Repository) *Service {
	return &Service{
		User:     user.NewUserService(repository.User),
		Products: products.NewProductsService(repository.Products),
	}
}
