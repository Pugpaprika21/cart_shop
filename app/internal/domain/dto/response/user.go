package response

type GetUsers struct {
	ID       *int64  `json:"id"`
	Username *string `json:"username"`
	Password *string `json:"password"`
	IsActive *int64  `json:"is_active"`
	CreBy    *string `json:"cre_by"`
	CreDate  *string `json:"cre_date"`
	UpdBy    *string `json:"upd_by"`
	UpdDate  *string `json:"upd_date"`
	ProgID   *string `json:"prog_id"`
	TotalRow *int64  `json:"total_row"`
}

type FindUser struct {
	ID       *int64  `json:"id"`
	Username *string `json:"username"`
	Password *string `json:"password"`
	IsActive *int64  `json:"is_active"`
	CreBy    *string `json:"cre_by"`
	CreDate  *string `json:"cre_date"`
	UpdBy    *string `json:"upd_by"`
	UpdDate  *string `json:"upd_date"`
	ProgID   *string `json:"prog_id"`
}
