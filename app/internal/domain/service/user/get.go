package user

import (
	"context"
	"miniservice/app/internal/domain/dto/qryparam"
	"miniservice/app/internal/domain/dto/request"
	"miniservice/app/internal/domain/dto/response"
	"miniservice/app/pkg/sqlx"
	"strings"
)

func (u *userService) GetUsers(ctx context.Context, req *request.GetUsers, qry *qryparam.GetUsers) ([]response.GetUsers, int64, error) {
	var sqlstr strings.Builder
	var whereClauses []string
	var args []interface{}
	var totalRow int64

	sqlstr.WriteString(`
		SELECT id, username, password, cre_by, cre_date, upd_by, upd_date, prog_id, is_active,
		(
			SELECT COUNT(1) FROM users WHERE is_active IS NULL OR is_active <> 1
		) AS total_row
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

	whereClauses = append(whereClauses, "is_active = ?")
	args = append(args, 0)

	if len(whereClauses) > 0 {
		sqlstr.WriteString(" WHERE ")
		sqlstr.WriteString(strings.Join(whereClauses, " AND "))
	}

	sqlstr.WriteString(" ORDER BY cre_date DESC")

	if req.Lazyload != nil && req.Lazyload.PageNo != nil && *req.Lazyload.PageNo != 0 && req.Lazyload.PageSize != nil && *req.Lazyload.PageSize != 0 {
		limit := *req.Lazyload.PageSize
		offset := (*req.Lazyload.PageNo - 1) * limit
		sqlstr.WriteString(" LIMIT ? OFFSET ?")
		args = append(args, limit, offset)
	}

	rows, err := u.repository.GetUsers(ctx, sqlx.Sqlx{Stmt: sqlstr.String(), Args: args})
	if err != nil {
		return nil, totalRow, err
	}

	resp := make([]response.GetUsers, len(rows))
	for i, rec := range rows {
		resp[i] = response.GetUsers{
			ID:       &rec.ID.Int64,
			Username: &rec.Username.String,
			Password: &rec.Password.String,
			IsActive: &rec.IsActive.Int64,
			CreBy:    &rec.CreBy.String,
			CreDate:  &rec.CreDate.String,
			UpdBy:    &rec.UpdBy.String,
			UpdDate:  &rec.UpdDate.String,
			ProgID:   &rec.ProgID.String,
			TotalRow: &rec.TotalRow.Int64,
		}
	}

	if len(resp) > 0 {
		totalRow = *resp[0].TotalRow
	}

	return resp, totalRow, nil
}
