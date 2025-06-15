package user

import (
	"context"
	"errors"
	"miniservice/app/internal/domain/dto/request"
	"miniservice/app/internal/enum"
	"miniservice/app/internal/persistence/schema"
	"miniservice/app/pkg/sqlx"
	"strings"
)

func (u *userService) UpdUser(ctx context.Context, req *request.UpdUser) error {
	var sqlstr strings.Builder
	var whereClauses []string
	var args []interface{}
	var dt string = sqlx.DateTimeNow()
	var rows []*request.UpdUserRows = req.UpdUserRows

	if rows == nil || len(rows) == 0 {
		return errors.New(enum.NO_DATA_PROVIDED_UPD_STR)
	}

	for _, rec := range rows {
		whereClauses = []string{}
		args = []interface{}{}

		params := schema.UpdUser{
			Username: sqlx.Nil(rec.Username),
			Password: sqlx.Nil(rec.Password),
			Email:    sqlx.Nil(rec.Email),
			IsActive: sqlx.Nil(rec.IsActive),
			UpdBy:    sqlx.Nil(rec.UpdBy),
			UpdDate:  &dt,
			ProgID:   rec.ProgID,
		}

		if rec.ID != nil && *rec.ID != 0 {
			whereClauses = append(whereClauses, "id = ?")
			args = append(args, rec.ID)
		}

		if len(whereClauses) > 0 {
			sqlstr.WriteString(strings.Join(whereClauses, " AND "))
		}

		err := u.repository.UpdUser(ctx, params, sqlx.Sqlx{Stmt: sqlstr.String(), Args: args})
		if err != nil {
			return err
		}

		sqlstr.Reset()
	}

	return nil
}
