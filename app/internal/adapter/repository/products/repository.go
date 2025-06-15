package products

import (
	"context"
	"miniservice/app/internal/persistence/schema"
	"miniservice/app/pkg/sqlx"

	"gorm.io/gorm"
)

type IProductsRepository interface {
	GenSKU(ctx context.Context, sql sqlx.Sqlx) ([]schema.GenSKU, error)
	CreProducts(ctx context.Context, params []schema.CreProducts) error
	GetProducts(ctx context.Context, sql sqlx.Sqlx) ([]schema.GetProducts, error)
	FindProducts(ctx context.Context, sql sqlx.Sqlx) (*schema.FindProducts, error)
	UpdProducts(ctx context.Context, params schema.UpdProducts, sql sqlx.Sqlx) error
	DelProducts(ctx context.Context, sql sqlx.Sqlx) error
	DelProductsIsActive(ctx context.Context, sql sqlx.Sqlx) error
}

type productsRepository struct {
	db *gorm.DB
}

func NewProductsRepository(db *gorm.DB) IProductsRepository {
	return &productsRepository{db: db}
}
