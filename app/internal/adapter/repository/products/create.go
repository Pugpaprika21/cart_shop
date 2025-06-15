package products

import (
	"context"
	"miniservice/app/internal/persistence/schema"
	"miniservice/app/pkg/errors"
)

func (p *productsRepository) CreProducts(ctx context.Context, params []schema.CreProducts) error {
	tx := p.db.WithContext(ctx).Begin()

	if err := tx.Error; err != nil {
		return errors.WrapDBError(err)
	}

	if err := tx.Table("products").Create(&params).Error; err != nil {
		tx.Rollback()
		return errors.WrapDBError(err)
	}

	if err := tx.Commit().Error; err != nil {
		return errors.WrapDBError(err)
	}

	return nil
}
