package router

import (
	"miniservice/app/internal/adapter/custmiddleware"
	"miniservice/app/internal/adapter/handler"
	"miniservice/app/pkg/jwtx"
	"net/http"
	"os"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ConfigVariantRouter struct {
	Handler *handler.Handler
	Jwtx    jwtx.IJwtx
}

type echoValidator struct {
	validator *validator.Validate
}

func (cv *echoValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

type router struct {
	server  *echo.Echo
	handler *handler.Handler
	jwtx    jwtx.IJwtx
}

func New(conf *ConfigVariantRouter) *router {
	return &router{
		server:  echo.New(),
		handler: conf.Handler,
		jwtx:    conf.Jwtx,
	}
}

func (r *router) Register() {
	route := r.server

	route.Use(custmiddleware.ReqLogger())
	route.Use(middleware.Recover())
	route.Use(middleware.CORS())
	route.Use(middleware.RequestID())
	route.Validator = &echoValidator{validator: validator.New()}

	r.v1()

	route.Logger.Fatal(route.Start(":" + os.Getenv("PORT_RUN")))
}
