package products

import (
	"context"
	"miniservice/app/internal/domain/dto/qryparam"
	"miniservice/app/internal/domain/dto/request"
	"miniservice/app/internal/domain/dto/response"
	"miniservice/app/pkg/sqlx"
	"strings"
)

func (p *productsService) GetProducts(ctx context.Context, req *request.GetProducts, qry *qryparam.GetProducts) ([]response.GetProducts, int64, error) {
	var sqlstr strings.Builder
	var whereClauses []string
	var args []interface{}
	var totalRow int64

	sqlstr.WriteString(`
		WITH total_products AS (
			SELECT COUNT(*) AS total FROM products WHERE is_active = 0
		)
		SELECT 
			p.id, 
			p.sku, 
			p.name, 
			p.description, 
			p.price, 
			p.stock, 
			p.is_active,
			tp.total AS total_row
		FROM products p
		CROSS JOIN total_products tp 
	`)

	if req.ID != nil && *req.ID != 0 {
		whereClauses = append(whereClauses, " p.id = ?")
		args = append(args, req.ID)
	}

	if req.SKU != nil && *req.SKU != "" {
		whereClauses = append(whereClauses, " p.sku = ?")
		args = append(args, req.SKU)
	}

	if req.Name != nil && *req.Name != "" {
		whereClauses = append(whereClauses, " p.name LIKE ?")
		args = append(args, "%"+strings.TrimSpace(*req.Name)+"%")
	}

	if req.Description != nil && *req.Description != "" {
		whereClauses = append(whereClauses, " p.description LIKE ?")
		args = append(args, "%"+strings.TrimSpace(*req.Description)+"%")
	}

	whereClauses = append(whereClauses, " p.is_active = ?")
	args = append(args, 0)

	if len(whereClauses) > 0 {
		sqlstr.WriteString(" WHERE ")
		sqlstr.WriteString(strings.Join(whereClauses, " AND "))
	}

	sqlstr.WriteString(" ORDER BY p.cre_date DESC")

	if req.Lazyload != nil && req.Lazyload.PageNo != nil && *req.Lazyload.PageNo != 0 && req.Lazyload.PageSize != nil && *req.Lazyload.PageSize != 0 {
		limit := *req.Lazyload.PageSize
		offset := (*req.Lazyload.PageNo - 1) * limit
		sqlstr.WriteString(" LIMIT ? OFFSET ?")
		args = append(args, limit, offset)
	}

	rows, err := p.repository.GetProducts(ctx, sqlx.Sqlx{Stmt: sqlstr.String(), Args: args})
	if err != nil {
		return nil, totalRow, err
	}

	resp := make([]response.GetProducts, len(rows))
	for i, rec := range rows {
		resp[i] = response.GetProducts{
			ID:          &rec.ID.Int64,
			SKU:         &rec.SKU.String,
			Name:        &rec.Name.String,
			Description: &rec.Description.String,
			Price:       &rec.Price.Float64,
			Stock:       &rec.Stock.Int64,
			IsActive:    &rec.IsActive.Int32,
			TotalRow:    &rec.TotalRow.Int64,
		}
	}

	if len(resp) > 0 {
		totalRow = *resp[0].TotalRow
	}

	return resp, totalRow, nil
}
