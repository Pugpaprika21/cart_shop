package schema

import "database/sql"

type GetProducts struct {
	ID          sql.NullInt64   `gorm:"column:id"`
	SKU         sql.NullString  `gorm:"column:sku"`
	Name        sql.NullString  `gorm:"column:name"`
	Description sql.NullString  `gorm:"column:description"`
	Price       sql.NullFloat64 `gorm:"column:price"`
	Stock       sql.NullInt64   `gorm:"column:stock"`
	CreBy       sql.NullString  `gorm:"column:cre_by"`
	CreDate     sql.NullString  `gorm:"column:cre_date"`
	UpdBy       sql.NullString  `gorm:"column:upd_by"`
	UpdDate     sql.NullString  `gorm:"column:upd_date"`
	ProgID      sql.NullString  `gorm:"column:prog_id"`
	IsActive    sql.NullInt32   `gorm:"column:is_active"`
	TotalRow    sql.NullInt64   `gorm:"column:total_row"`
}

type GetHasProductExisting struct {
	ProductExisting sql.NullInt64 `gorm:"column:product_existing"`
}

type FindProducts struct {
	ID          sql.NullInt64   `gorm:"column:id"`
	SKU         sql.NullString  `gorm:"column:sku"`
	Name        sql.NullString  `gorm:"column:name"`
	Description sql.NullString  `gorm:"column:description"`
	Price       sql.NullFloat64 `gorm:"column:price"`
	Stock       sql.NullInt64   `gorm:"column:stock"`
	CreBy       sql.NullString  `gorm:"column:cre_by"`
	CreDate     sql.NullString  `gorm:"column:cre_date"`
	UpdBy       sql.NullString  `gorm:"column:upd_by"`
	UpdDate     sql.NullString  `gorm:"column:upd_date"`
	ProgID      sql.NullString  `gorm:"column:prog_id"`
	IsActive    sql.NullInt32   `gorm:"column:is_active"`
	TotalRow    sql.NullInt64   `gorm:"column:total_row"`
}

type CreProducts struct {
	ID          *int64   `gorm:"column:id"`
	SKU         *string  `gorm:"column:sku"`
	Name        *string  `gorm:"column:name"`
	Description *string  `gorm:"column:description"`
	Price       *float64 `gorm:"column:price"`
	Stock       *int64   `gorm:"column:stock"`
	CreBy       *string  `gorm:"column:cre_by"`
	CreDate     *string  `gorm:"column:cre_date"`
	UpdBy       *string  `gorm:"column:upd_by"`
	UpdDate     *string  `gorm:"column:upd_date"`
	ProgID      *string  `gorm:"column:prog_id"`
	IsActive    *int32   `gorm:"column:is_active"`
}

type UpdProducts struct {
	ID          *int64   `gorm:"column:id"`
	SKU         *string  `gorm:"column:sku"`
	Name        *string  `gorm:"column:name"`
	Description *string  `gorm:"column:description"`
	Price       *float64 `gorm:"column:price"`
	Stock       *int64   `gorm:"column:stock"`
	CreBy       *string  `gorm:"column:cre_by"`
	CreDate     *string  `gorm:"column:cre_date"`
	UpdBy       *string  `gorm:"column:upd_by"`
	UpdDate     *string  `gorm:"column:upd_date"`
	ProgID      *string  `gorm:"column:prog_id"`
	IsActive    *int32   `gorm:"column:is_active"`
}

type GenSKU struct {
	ResultSKU sql.NullString `gorm:"column:result_sku"`
	MaxNum    sql.NullInt64  `gorm:"column:max_num"`
}
