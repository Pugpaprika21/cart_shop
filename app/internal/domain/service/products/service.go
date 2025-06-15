package products

import (
	"context"
	"miniservice/app/internal/adapter/repository/products"
	"miniservice/app/internal/domain/dto/qryparam"
	"miniservice/app/internal/domain/dto/request"
	"miniservice/app/internal/domain/dto/response"
)

type IProductsService interface {
	CreProducts(ctx context.Context, req *request.CreProducts) error
	GetProducts(ctx context.Context, req *request.GetProducts, qry *qryparam.GetProducts) ([]response.GetProducts, int64, error)
	FindProducts(ctx context.Context, req *request.FindProducts, qry *qryparam.FindProducts) (*response.FindProducts, error)
	UpdProducts(ctx context.Context, req *request.UpdProducts) error
	DelProducts(ctx context.Context, req *request.DelProducts) error
}

type productsService struct {
	repository products.IProductsRepository
}

func NewProductsService(repository products.IProductsRepository) IProductsService {
	return &productsService{repository: repository}
}
