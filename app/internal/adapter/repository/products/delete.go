package products

import (
	"context"
	"fmt"
	"miniservice/app/pkg/errors"
	"miniservice/app/pkg/sqlx"
)

func (p *productsRepository) DelProducts(ctx context.Context, sql sqlx.Sqlx) error {
	tx := p.db.WithContext(ctx).Begin()

	if err := tx.Error; err != nil {
		return errors.WrapDBError(err)
	}

	result := tx.Exec(sql.Stmt, sql.Args...)
	if err := result.Error; err != nil {
		tx.Rollback()
		return errors.WrapDBError(err)
	}

	if rowsAffected := result.RowsAffected; rowsAffected == 0 {
		tx.Rollback()
		return fmt.Errorf("No Records Deleted")
	}

	if err := tx.Commit().Error; err != nil {
		return errors.WrapDBError(err)
	}

	return nil
}

func (p *productsRepository) DelProductsIsActive(ctx context.Context, sql sqlx.Sqlx) error {
	tx := p.db.WithContext(ctx).Begin()

	if err := tx.Error; err != nil {
		return errors.WrapDBError(err)
	}

	result := tx.Exec(sql.Stmt, sql.Args...)
	if err := result.Error; err != nil {
		tx.Rollback()
		return errors.WrapDBError(err)
	}

	if err := tx.Commit().Error; err != nil {
		return errors.WrapDBError(err)
	}

	return nil
}
