package user

import (
	"context"
	"miniservice/app/internal/adapter/repository/user"
	"miniservice/app/internal/domain/dto/qryparam"
	"miniservice/app/internal/domain/dto/request"
	"miniservice/app/internal/domain/dto/response"
)

type IUserService interface {
	GetUsers(ctx context.Context, req *request.GetUsers, qry *qryparam.GetUsers) ([]response.GetUsers, int64, error)
	FindUser(ctx context.Context, req *request.FindUser, qry *qryparam.FindUser) (*response.FindUser, error)
	CreUsers(ctx context.Context, req *request.CreUsers) error
	UpdUser(ctx context.Context, req *request.UpdUser) error
	DelUser(ctx context.Context, req *request.DelUser) error
}

type userService struct {
	repository user.IUserRepository
}

func NewUserService(repository user.IUserRepository) IUserService {
	return &userService{repository: repository}
}
