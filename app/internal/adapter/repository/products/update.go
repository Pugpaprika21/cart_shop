package products

import (
	"context"
	"miniservice/app/internal/persistence/schema"
	"miniservice/app/pkg/errors"
	"miniservice/app/pkg/sqlx"
)

func (p *productsRepository) UpdProducts(ctx context.Context, params schema.UpdProducts, sql sqlx.Sqlx) error {
	tx := p.db.WithContext(ctx).Begin()

	if err := tx.Error; err != nil {
		return errors.WrapDBError(err)
	}

	if err := tx.Table("products").Where(sql.Stmt, sql.Args...).Updates(&params).Error; err != nil {
		tx.Rollback()
		return errors.WrapDBError(err)
	}

	if err := tx.Commit().Error; err != nil {
		return errors.WrapDBError(err)
	}

	return nil
}
