package products

import (
	"miniservice/app/internal/domain/service/products"

	"github.com/labstack/echo/v4"
)

type IProductsHandler interface {
	CreProducts(c echo.Context) error
}

type productsHandler struct {
	service products.IProductsService
}

func NewProductsHandler(service products.IProductsService) IProductsHandler {
	return &productsHandler{service: service}
}
