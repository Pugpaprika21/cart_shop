package products

import (
	"context"
	"errors"
	"fmt"
	"miniservice/app/internal/domain/dto/request"
	"miniservice/app/internal/enum"
	"miniservice/app/internal/persistence/schema"
	"miniservice/app/pkg/sqlx"
	"strings"
	"time"
)

func (p *productsService) CreProducts(ctx context.Context, req *request.CreProducts) error {
	var sqlstr strings.Builder
	dt := sqlx.DateTimeNow()

	var rows []*request.CreProductsRows = req.CreProductsRows
	if rows == nil || len(rows) == 0 {
		return errors.New(enum.NO_DATA_PROVIDED_CRE_STR)
	}

	params := make([]schema.CreProducts, len(rows))
	prefix := time.Now().Format("0601")

	sqlstr.WriteString(`
		SELECT IFNULL(MAX(CAST(SUBSTRING(sku, 5, 4) AS UNSIGNED)), 0) AS max_num
		FROM products
		WHERE LEFT(sku, 4) = ?;
	`)

	genSku, err := p.repository.GenSKU(ctx, sqlx.Sqlx{Stmt: sqlstr.String(), Args: []interface{}{prefix}})
	if err != nil {
		return err
	}

	currentNum := 0
	if len(genSku) > 0 {
		currentNum = int(genSku[0].MaxNum.Int64)
	}

	sqlstr.Reset()

	for i, rec := range rows {
		if rec.SKU == nil || *rec.SKU == "" {
			currentNum++
			sku := fmt.Sprintf("%s%04d", prefix, currentNum)
			rec.SKU = &sku
		} else {
			sqlstr.WriteString(`SELECT COUNT(*) AS product_existing FROM products WHERE sku = ? AND is_active = 0`)

			checkRes, err := p.repository.GetHasProductExisting(ctx, sqlx.Sqlx{Stmt: sqlstr.String(), Args: []interface{}{*rec.SKU}})
			if err != nil {
				return err
			}
			sqlstr.Reset()

			if len(checkRes) > 0 && checkRes[0].ProductExisting.Int64 > 0 {
				return fmt.Errorf("SKU '%s' ถูกใช้แล้ว", *rec.SKU)
			}
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
