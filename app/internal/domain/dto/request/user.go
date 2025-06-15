package request

type GetUsers struct {
	Lazyload *Lazyload `json:"lazyload"`
	ID       *int64    `json:"id"`
	Username *string   `json:"username"`
	Password *string   `json:"password"`
	Email    *string   `json:"email"`
	IsActive *int32    `json:"is_active"`
	CreBy    *string   `json:"cre_by"`
	CreDate  *string   `json:"cre_date"`
	UpdBy    *string   `json:"upd_by"`
	UpdDate  *string   `json:"upd_date"`
	ProgID   *string   `json:"prog_id"`
}

type FindUser struct {
	ID       *int64  `json:"id" validate:"required"`
	Username *string `json:"username"`
	Password *string `json:"password"`
	Email    *string `json:"email"`
	IsActive *int32  `json:"is_active"`
	CreBy    *string `json:"cre_by"`
	CreDate  *string `json:"cre_date"`
	UpdBy    *string `json:"upd_by"`
	UpdDate  *string `json:"upd_date"`
	ProgID   *string `json:"prog_id"`
}

type CreUsersRows struct {
	ID       *int64  `json:"id"`
	Username *string `json:"username"`
	Password *string `json:"password"`
	Email    *string `json:"email"`
	IsActive *int32  `json:"is_active"`
	CreBy    *string `json:"cre_by"`
	CreDate  *string `json:"cre_date"`
	UpdBy    *string `json:"upd_by"`
	UpdDate  *string `json:"upd_date"`
	ProgID   *string `json:"prog_id"`
}

type CreUsers struct {
	CreUsersRows []*CreUsersRows `json:"cre_users_rows"`
}

type UpdUserRows struct {
	ID       *int64  `json:"id"`
	Username *string `json:"username"`
	Password *string `json:"password"`
	Email    *string `json:"email"`
	IsActive *int32  `json:"is_active"`
	CreBy    *string `json:"cre_by"`
	CreDate  *string `json:"cre_date"`
	UpdBy    *string `json:"upd_by"`
	UpdDate  *string `json:"upd_date"`
	ProgID   *string `json:"prog_id"`
}

type UpdUser struct {
	UpdUserRows []*UpdUserRows `json:"upd_user_rows"`
}

type DelUser struct {
	ID       *int64  `json:"id" validate:"required"`
	IsActive *string `json:"is_active"`
}

type DelUserIsActive struct {
	ID       *int64  `json:"id" validate:"required"`
	IsActive *string `json:"is_active"`
}
