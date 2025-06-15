package products

import (
	"context"
	"miniservice/app/internal/persistence/schema"
	"miniservice/app/pkg/errors"
	"miniservice/app/pkg/sqlx"
)

func (p *productsRepository) FindProducts(ctx context.Context, sql sqlx.Sqlx) (*schema.FindProducts, error) {
	var row schema.FindProducts
	result := p.db.WithContext(ctx).Raw(sql.Stmt, sql.Args...).First(&row)
	if result.Error != nil {
		return nil, errors.WrapDBError(result.Error)
	}

	return &row, nil
}
