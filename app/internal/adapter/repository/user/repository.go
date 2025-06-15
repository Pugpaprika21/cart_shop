package user

import (
	"context"
	"miniservice/app/internal/persistence/schema"
	"miniservice/app/pkg/sqlx"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUsers(ctx context.Context, sql sqlx.Sqlx) ([]schema.GetUsers, error)
	FindUser(ctx context.Context, sql sqlx.Sqlx) (*schema.FindUser, error)
	CreUsers(ctx context.Context, params []schema.CreUsers) error
	UpdUser(ctx context.Context, params schema.UpdUser, sql sqlx.Sqlx) error
	DelUser(ctx context.Context, sql sqlx.Sqlx) error
	DelUserIsActive(ctx context.Context, sql sqlx.Sqlx) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db: db}
}
