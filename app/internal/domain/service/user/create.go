package user

import (
	"context"
	"errors"
	"miniservice/app/internal/domain/dto/request"
	"miniservice/app/internal/enum"
	"miniservice/app/internal/persistence/schema"
	"miniservice/app/pkg/sqlx"
)

func (u *userService) CreUsers(ctx context.Context, req *request.CreUsers) error {
	var dt string = sqlx.DateTimeNow()
	var rows []*request.CreUsersRows = req.CreUsersRows
	if rows == nil || len(rows) == 0 {
		return errors.New(enum.NO_DATA_PROVIDED_CRE_STR)
	}

	params := make([]schema.CreUsers, len(rows))
	for i, rec := range rows {
		params[i] = schema.CreUsers{
			Username: sqlx.Nil(rec.Username),
			Password: sqlx.Nil(rec.Password),
			Email:    sqlx.Nil(rec.Email),
			IsActive: sqlx.Nil(rec.IsActive),
			CreBy:    sqlx.Nil(rec.CreBy),
			CreDate:  &dt,
			UpdBy:    sqlx.Nil(rec.UpdBy),
			UpdDate:  sqlx.Nil(rec.UpdDate),
			ProgID:   sqlx.Nil(rec.ProgID),
		}
	}

	err := u.repository.CreUsers(ctx, params)
	if err != nil {
		return err
	}

	return nil
}
