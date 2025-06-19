package products

import (
	"log/slog"
	"miniservice/app/internal/domain/dto/request"
	"miniservice/app/internal/enum"
	"miniservice/app/pkg/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (p *productsHandler) DelProducts(c echo.Context) error {
	var req request.DelProducts
	var reqID = c.Response().Header().Get(echo.HeaderXRequestID)
	var resp = response.NewResponseBuilder()

	if err := c.Bind(&req); err != nil {
		p.logger.Error("bind error", slog.Any("handler", req))
		return c.JSON(
			http.StatusBadRequest,
			resp.Message(err.Error()).Code(enum.FOR_BAD_REQUEST).RequestID(reqID).Build(),
		)
	}

	if err := c.Validate(&req); err != nil {
		p.logger.Error("validate error", slog.Any("handler", req))
		return c.JSON(
			http.StatusBadRequest,
			resp.Message(err.Error()).Code(enum.FOR_BAD_REQUEST).RequestID(reqID).Build(),
		)
	}

	err := p.service.DelProducts(c.Request().Context(), &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resp.Message(err.Error()).Code(enum.FOR_ERROR).RequestID(reqID).Build())
	}

	return c.JSON(http.StatusOK, resp.Message(enum.SUCCESS_STR).Code(200).RequestID(reqID).Build())
}
