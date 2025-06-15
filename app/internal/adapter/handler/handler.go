package handler

import (
	"miniservice/app/internal/adapter/handler/products"
	"miniservice/app/internal/adapter/handler/user"
	"miniservice/app/internal/domain/service"
)

type Handler struct {
	User     user.IUserHandler
	Products products.IProductsHandler
}

func New(service *service.Service) *Handler {
	return &Handler{
		User:     user.NewUserhandler(service.User),
		Products: products.NewProductsHandler(service.Products),
	}
}
