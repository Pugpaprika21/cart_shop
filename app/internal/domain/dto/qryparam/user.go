package qryparam

type GetUsers struct {
	ID       *int64  `query:"id"`
	Username *string `query:"username"`
	Password *string `query:"password"`
	Email    *string `query:"email"`
	IsActive *int32  `query:"is_active"`
	CreBy    *string `query:"cre_by"`
	CreDate  *string `query:"cre_date"`
	UpdBy    *string `query:"upd_by"`
	UpdDate  *string `query:"upd_date"`
	ProgID   *string `query:"prog_id"`
}

type FindUser struct {
	ID       *int64  `query:"id"`
	Username *string `query:"username"`
	Password *string `query:"password"`
	Email    *string `query:"email"`
	IsActive *int32  `query:"is_active"`
	CreBy    *string `query:"cre_by"`
	CreDate  *string `query:"cre_date"`
	UpdBy    *string `query:"upd_by"`
	UpdDate  *string `query:"upd_date"`
	ProgID   *string `query:"prog_id"`
}
