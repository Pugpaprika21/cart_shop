package qryparam

type GetProducts struct {
	ID          *int64   `query:"id"`
	SKU         *string  `query:"sku"`
	Name        *string  `query:"name"`
	Description *string  `query:"description"`
	Price       *float64 `query:"price"`
	Stock       *int64   `query:"stock"`
	CreBy       *string  `query:"cre_by"`
	CreDate     *string  `query:"cre_date"`
	UpdBy       *string  `query:"upd_by"`
	UpdDate     *string  `query:"upd_date"`
	ProgID      *string  `query:"prog_id"`
	IsActive    *int32   `query:"is_active"`
	TotalRow    *int64   `query:"total_row"`
}

type FindProducts struct {
	ID          *int64   `query:"id" validate:"required"`
	SKU         *string  `query:"sku"`
	Name        *string  `query:"name"`
	Description *string  `query:"description"`
	Price       *float64 `query:"price"`
	Stock       *int64   `query:"stock"`
	CreBy       *string  `query:"cre_by"`
	CreDate     *string  `query:"cre_date"`
	UpdBy       *string  `query:"upd_by"`
	UpdDate     *string  `query:"upd_date"`
	ProgID      *string  `query:"prog_id"`
	IsActive    *int32   `query:"is_active"`
	TotalRow    *int64   `query:"total_row"`
}
