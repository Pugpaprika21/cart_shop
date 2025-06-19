package products

import (
	"log/slog"
	"miniservice/app/internal/domain/service/products"
	"os"

	"github.com/labstack/echo/v4"
)

type IProductsHandler interface {
	CreProducts(c echo.Context) error
	GetProducts(c echo.Context) error
	FindProducts(c echo.Context) error
	UpdProducts(c echo.Context) error
	DelProducts(c echo.Context) error
}

type productsHandler struct {
	service products.IProductsService
	logger  *slog.Logger
}

func NewProductsHandler(service products.IProductsService) IProductsHandler {
	return &productsHandler{service: service, logger: slog.New(slog.NewJSONHandler(os.Stdout, nil))}
}
