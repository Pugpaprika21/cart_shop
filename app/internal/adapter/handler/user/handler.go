package user

import (
	"log/slog"
	"miniservice/app/internal/domain/service/user"
	"os"

	"github.com/labstack/echo/v4"
)

type IUserHandler interface {
	GetUsers(c echo.Context) error
	FindUser(c echo.Context) error
	CreUsers(c echo.Context) error
	UpdUser(c echo.Context) error
	DelUser(c echo.Context) error
}

type userHandler struct {
	service user.IUserService
	logger  *slog.Logger
}

func NewUserhandler(service user.IUserService) IUserHandler {
	return &userHandler{service: service, logger: slog.New(slog.NewJSONHandler(os.Stdout, nil))}
}
