package user

import (
	"context"
	"miniservice/app/internal/domain/dto/qryparam"
	"miniservice/app/internal/domain/dto/request"
	"miniservice/app/internal/domain/dto/response"
	"miniservice/app/pkg/sqlx"
	"strings"
)

func (u *userService) FindUser(ctx context.Context, req *request.FindUser, qry *qryparam.FindUser) (*response.FindUser, error) {
	var sqlstr strings.Builder
	var whereClauses []string
	var args []interface{}

	sqlstr.WriteString(`
		SELECT id, username, password, cre_by, cre_date, upd_by, upd_date, prog_id, is_active
		FROM users
	`)

	if req.ID != nil && *req.ID != 0 {
		whereClauses = append(whereClauses, "id = ?")
		args = append(args, req.ID)
	}

	if req.Username != nil && *req.Username != "" {
		whereClauses = append(whereClauses, "username = ?")
		args = append(args, req.Username)
	}

	if req.Password != nil && *req.Password != "" {
		whereClauses = append(whereClauses, "password = ?")
		args = append(args, req.Password)
	}

	if len(whereClauses) > 0 {
		sqlstr.WriteString(" WHERE ")
		sqlstr.WriteString(strings.Join(whereClauses, " AND "))
	}

	rec, err := u.repository.FindUser(ctx, sqlx.Sqlx{Stmt: sqlstr.String(), Args: args})
	if err != nil {
		return nil, err
	}

	resp := response.FindUser{
		ID:       &rec.ID.Int64,
		Username: &rec.Username.String,
		Password: &rec.Password.String,
		IsActive: &rec.IsActive.Int64,
		CreBy:    &rec.CreBy.String,
		CreDate:  &rec.CreDate.String,
		UpdBy:    &rec.UpdBy.String,
		UpdDate:  &rec.UpdDate.String,
		ProgID:   &rec.ProgID.String,
	}

	return &resp, nil
}
