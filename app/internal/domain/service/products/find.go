package products

import (
	"context"
	"miniservice/app/internal/domain/dto/qryparam"
	"miniservice/app/internal/domain/dto/request"
	"miniservice/app/internal/domain/dto/response"
	"miniservice/app/pkg/sqlx"
	"strings"
)

func (p *productsService) FindProducts(ctx context.Context, req *request.FindProducts, qry *qryparam.FindProducts) (*response.FindProducts, error) {
	var sqlstr strings.Builder
	var whereClauses []string
	var args []interface{}

	sqlstr.WriteString(`SELECT id, sku, name, description, price, stock, is_active FROM products p`)

	if req.ID != nil && *req.ID != 0 {
		whereClauses = append(whereClauses, " id = ?")
		args = append(args, req.ID)
	}

	if req.SKU != nil && *req.SKU != "" {
		whereClauses = append(whereClauses, " sku = ?")
		args = append(args, req.SKU)
	}

	if req.Name != nil && *req.Name != "" {
		whereClauses = append(whereClauses, " name LIKE ?")
		args = append(args, "%"+strings.TrimSpace(*req.Name)+"%")
	}

	if req.Description != nil && *req.Description != "" {
		whereClauses = append(whereClauses, " description LIKE ?")
		args = append(args, "%"+strings.TrimSpace(*req.Description)+"%")
	}

	whereClauses = append(whereClauses, " is_active = ?")
	args = append(args, 0)

	if len(whereClauses) > 0 {
		sqlstr.WriteString(" WHERE ")
		sqlstr.WriteString(strings.Join(whereClauses, " AND "))
	}

	sqlstr.WriteString(" ORDER BY cre_date DESC")

	rec, err := p.repository.FindProducts(ctx, sqlx.Sqlx{Stmt: sqlstr.String(), Args: args})
	if err != nil {
		return nil, err
	}

	resp := response.FindProducts{
		ID:          &rec.ID.Int64,
		SKU:         &rec.SKU.String,
		Name:        &rec.Name.String,
		Description: &rec.Description.String,
		IsActive:    &rec.IsActive.Int32,
		CreBy:       &rec.CreBy.String,
		CreDate:     &rec.CreDate.String,
		UpdBy:       &rec.UpdBy.String,
		UpdDate:     &rec.UpdDate.String,
		ProgID:      &rec.ProgID.String,
	}

	return &resp, nil
}
