package user

import (
	"context"
	"miniservice/app/internal/persistence/schema"
	"miniservice/app/pkg/errors"
	"miniservice/app/pkg/sqlx"
)

func (u *userRepository) GetUsers(ctx context.Context, sql sqlx.Sqlx) ([]schema.GetUsers, error) {
	var rows []schema.GetUsers
	result := u.db.WithContext(ctx).Raw(sql.Stmt, sql.Args...).Scan(&rows)
	if result.Error != nil {
		return nil, errors.WrapDBError(result.Error)
	}

	return rows, nil
}
