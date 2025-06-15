package schema

import "database/sql"

type GetUsers struct {
	ID       sql.NullInt64  `gorm:"column:id"`
	Username sql.NullString `gorm:"column:username"`
	Password sql.NullString `gorm:"column:password"`
	Email    sql.NullString `gorm:"column:email"`
	CreBy    sql.NullString `gorm:"column:cre_by"`
	CreDate  sql.NullString `gorm:"column:cre_date"`
	UpdBy    sql.NullString `gorm:"column:upd_by"`
	UpdDate  sql.NullString `gorm:"column:upd_date"`
	ProgID   sql.NullString `gorm:"column:prog_id"`
	IsActive sql.NullInt64  `gorm:"column:is_active"`
	TotalRow sql.NullInt64  `gorm:"column:total_row"`
}

type FindUser struct {
	ID       sql.NullInt64  `gorm:"column:id"`
	Username sql.NullString `gorm:"column:username"`
	Password sql.NullString `gorm:"column:password"`
	Email    sql.NullString `gorm:"column:email"`
	CreBy    sql.NullString `gorm:"column:cre_by"`
	CreDate  sql.NullString `gorm:"column:cre_date"`
	UpdBy    sql.NullString `gorm:"column:upd_by"`
	UpdDate  sql.NullString `gorm:"column:upd_date"`
	ProgID   sql.NullString `gorm:"column:prog_id"`
	IsActive sql.NullInt64  `gorm:"column:is_active"`
}

type CreUsers struct {
	ID       *int64  `gorm:"column:id"`
	Username *string `gorm:"column:username"`
	Password *string `gorm:"column:password"`
	Email    *string `gorm:"column:email"`
	CreBy    *string `gorm:"column:cre_by"`
	CreDate  *string `gorm:"column:cre_date"`
	UpdBy    *string `gorm:"column:upd_by"`
	UpdDate  *string `gorm:"column:upd_date"`
	ProgID   *string `gorm:"column:prog_id"`
	IsActive *int32  `gorm:"column:is_active"`
}

type UpdUser struct {
	ID       *int64  `gorm:"column:id"`
	Username *string `gorm:"column:username"`
	Password *string `gorm:"column:password"`
	Email    *string `gorm:"column:email"`
	CreBy    *string `gorm:"column:cre_by"`
	CreDate  *string `gorm:"column:cre_date"`
	UpdBy    *string `gorm:"column:upd_by"`
	UpdDate  *string `gorm:"column:upd_date"`
	ProgID   *string `gorm:"column:prog_id"`
	IsActive *int32  `gorm:"column:is_active"`
}
