package request

type CreProductsRows struct {
	ID          *int64   `json:"id"`
	SKU         *string  `json:"sku"`
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Price       *float64 `json:"price"`
	Stock       *int64   `json:"stock"`
	CreBy       *string  `json:"cre_by"`
	CreDate     *string  `json:"cre_date"`
	UpdBy       *string  `json:"upd_by"`
	UpdDate     *string  `json:"upd_date"`
	ProgID      *string  `json:"prog_id"`
	IsActive    *int32   `json:"is_active"`
	TotalRow    *int64   `json:"total_row"`
}

type CreProducts struct {
	CreProductsRows []*CreProductsRows `json:"cre_products_rows"`
}

type UpdProductsRows struct {
	ID          *int64   `json:"id"`
	SKU         *string  `json:"sku"`
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Price       *float64 `json:"price"`
	Stock       *int64   `json:"stock"`
	CreBy       *string  `json:"cre_by"`
	CreDate     *string  `json:"cre_date"`
	UpdBy       *string  `json:"upd_by"`
	UpdDate     *string  `json:"upd_date"`
	ProgID      *string  `json:"prog_id"`
	IsActive    *int32   `json:"is_active"`
	TotalRow    *int64   `json:"total_row"`
}

type UpdProducts struct {
	UpdProductsRows []*UpdProductsRows `json:"upd_products_rows"`
}

type GetProducts struct {
	Lazyload    *Lazyload `json:"lazyload"`
	ID          *int64    `json:"id"`
	SKU         *string   `json:"sku"`
	Name        *string   `json:"name"`
	Description *string   `json:"description"`
	Price       *float64  `json:"price"`
	Stock       *int64    `json:"stock"`
	CreBy       *string   `json:"cre_by"`
	CreDate     *string   `json:"cre_date"`
	UpdBy       *string   `json:"upd_by"`
	UpdDate     *string   `json:"upd_date"`
	ProgID      *string   `json:"prog_id"`
	IsActive    *int32    `json:"is_active"`
	TotalRow    *int64    `json:"total_row"`
}

type FindProducts struct {
	ID          *int64   `json:"id"`
	SKU         *string  `json:"sku"`
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Price       *float64 `json:"price"`
	Stock       *int64   `json:"stock"`
	CreBy       *string  `json:"cre_by"`
	CreDate     *string  `json:"cre_date"`
	UpdBy       *string  `json:"upd_by"`
	UpdDate     *string  `json:"upd_date"`
	ProgID      *string  `json:"prog_id"`
	IsActive    *int32   `json:"is_active"`
	TotalRow    *int64   `json:"total_row"`
}

type DelProducts struct {
	ID       *int64  `json:"id" validate:"required"`
	IsActive *string `json:"is_active"`
}
