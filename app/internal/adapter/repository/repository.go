package repository

import (
	"miniservice/app/internal/adapter/repository/products"
	"miniservice/app/internal/adapter/repository/user"

	"gorm.io/gorm"
)

type Repository struct {
	User     user.IUserRepository
	Products products.IProductsRepository
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		User:     user.NewUserRepository(db),
		Products: products.NewProductsRepository(db),
	}
}
