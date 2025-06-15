package products

import (
	"context"
	"miniservice/app/internal/adapter/repository/products"
	"miniservice/app/internal/domain/dto/request"
)

type IProductsService interface {
	CreProducts(ctx context.Context, req *request.CreProducts) error
}

type productsService struct {
	repository products.IProductsRepository
}

func NewProductsService(repository products.IProductsRepository) IProductsService {
	return &productsService{repository: repository}
}
