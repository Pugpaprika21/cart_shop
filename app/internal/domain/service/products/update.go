package products

import (
	"context"
	"errors"
	"miniservice/app/internal/domain/dto/request"
	"miniservice/app/internal/enum"
	"miniservice/app/internal/persistence/schema"
	"miniservice/app/pkg/sqlx"
	"strings"
)

func (p *productsService) UpdProducts(ctx context.Context, req *request.UpdProducts) error {
	var sqlstr strings.Builder
	var whereClauses []string
	var args []interface{}
	var dt string = sqlx.DateTimeNow()
	var rows []*request.UpdProductsRows = req.UpdProductsRows

	if rows == nil || len(rows) == 0 {
		return errors.New(enum.NO_DATA_PROVIDED_UPD_STR)
	}

	for _, rec := range rows {
		whereClauses = []string{}
		args = []interface{}{}

		params := schema.UpdProducts{
			SKU:         rec.SKU,
			Name:        rec.Name,
			Description: rec.Description,
			Price:       rec.Price,
			Stock:       rec.Stock,
			CreBy:       rec.CreBy,
			CreDate:     &dt,
			UpdBy:       rec.UpdBy,
			UpdDate:     rec.UpdDate,
			ProgID:      rec.ProgID,
			IsActive:    rec.IsActive,
		}

		if rec.ID != nil && *rec.ID != 0 {
			whereClauses = append(whereClauses, "id = ?")
			args = append(args, rec.ID)
		}

		if len(whereClauses) > 0 {
			sqlstr.WriteString(strings.Join(whereClauses, " AND "))
		}

		err := p.repository.UpdProducts(ctx, params, sqlx.Sqlx{Stmt: sqlstr.String(), Args: args})
		if err != nil {
			return err
		}

		sqlstr.Reset()
	}

	return nil
}
