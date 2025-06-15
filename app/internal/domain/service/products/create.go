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

func (p *productsService) CreProducts(ctx context.Context, req *request.CreProducts) error {
	var sqlstr strings.Builder
	var args []interface{}
	dt := sqlx.DateTimeNow()

	rows := req.CreProductsRows
	if rows == nil || len(rows) == 0 {
		return errors.New(enum.NO_DATA_PROVIDED_CRE_STR)
	}

	params := make([]schema.CreProducts, len(rows))
	for i, rec := range rows {
		var resultSKU string

		if rec.SKU == nil || *rec.SKU == "" {
			sqlstr.Reset()
			sqlstr.WriteString(`
				SELECT CONCAT(DATE_FORMAT(NOW(), '%y%m'), LPAD(IFNULL(MAX(CAST(SUBSTRING(sku, 5, 4) AS UNSIGNED)) + 1, 1), 4, '0')) AS result_sku
				FROM products
				WHERE LEFT(sku, 4) = DATE_FORMAT(NOW(), '%y%m');
			`)

			genSku, err := p.repository.GenSKU(ctx, sqlx.Sqlx{Stmt: sqlstr.String(), Args: args})
			if err != nil {
				return err
			}

			resultSKU = genSku[0].ResultSKU.String
			rec.SKU = &resultSKU
		}

		params[i] = schema.CreProducts{
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
	}

	if err := p.repository.CreProducts(ctx, params); err != nil {
		return err
	}

	return nil
}
