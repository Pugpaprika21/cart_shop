package user

import (
	"context"
	"miniservice/app/internal/domain/dto/request"
	"miniservice/app/internal/enum"
	"miniservice/app/pkg/sqlx"
	"strings"
)

func (u *userService) DelUser(ctx context.Context, req *request.DelUser) error {
	var whereBuilder strings.Builder
	var args []interface{}
	var sqlstr strings.Builder
	var err error

	if req.ID != nil && *req.ID != 0 {
		whereBuilder.WriteString(" WHERE id = ?")
		args = append(args, req.ID)
	}

	if req.IsActive != nil && *req.IsActive == enum.IS_ACTIVE {
		sqlstr.WriteString(`UPDATE users SET upd_date = NOW(), is_active = '1'`)
		sqlstr.WriteString(whereBuilder.String())
		err = u.repository.DelUserIsActive(ctx, sqlx.Sqlx{Stmt: sqlstr.String(), Args: args})
		sqlstr.Reset()
	} else {
		sqlstr.WriteString(`DELETE FROM users`)
		sqlstr.WriteString(whereBuilder.String())
		err = u.repository.DelUser(ctx, sqlx.Sqlx{Stmt: sqlstr.String(), Args: args})
		sqlstr.Reset()
	}

	if err != nil {
		return err
	}

	return nil
}
