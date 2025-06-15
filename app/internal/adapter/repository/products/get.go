package products

import (
	"context"
	"miniservice/app/internal/persistence/schema"
	"miniservice/app/pkg/errors"
	"miniservice/app/pkg/sqlx"
)

func (p *productsRepository) GenSKU(ctx context.Context, sql sqlx.Sqlx) ([]schema.GenSKU, error) {
	var rows []schema.GenSKU
	result := p.db.WithContext(ctx).Raw(sql.Stmt, sql.Args...).Scan(&rows)
	if result.Error != nil {
		return nil, errors.WrapDBError(result.Error)
	}

	return rows, nil
}

func (p *productsRepository) GetProducts(ctx context.Context, sql sqlx.Sqlx) ([]schema.GetProducts, error) {
	var rows []schema.GetProducts
	result := p.db.WithContext(ctx).Raw(sql.Stmt, sql.Args...).Scan(&rows)
	if result.Error != nil {
		return nil, errors.WrapDBError(result.Error)
	}

	return rows, nil
}
