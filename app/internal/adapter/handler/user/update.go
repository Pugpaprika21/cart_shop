package user

import (
	"miniservice/app/internal/domain/dto/request"
	"miniservice/app/internal/enum"
	"miniservice/app/pkg/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (u *userHandler) UpdUser(c echo.Context) error {
	var req request.UpdUser
	var reqID = c.Response().Header().Get(echo.HeaderXRequestID)
	var resp = response.NewResponseBuilder()

	if err := c.Bind(&req); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			resp.Message(err.Error()).Code(enum.FOR_BAD_REQUEST).RequestID(reqID).Build(),
		)
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			resp.Message(err.Error()).Code(enum.FOR_BAD_REQUEST).RequestID(reqID).Build(),
		)
	}

	err := u.service.UpdUser(c.Request().Context(), &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resp.Message(err.Error()).Code(enum.FOR_ERROR).RequestID(reqID).Build())
	}

	return c.JSON(http.StatusOK, resp.Message(enum.SUCCESS_STR).Code(200).RequestID(reqID).Build())
}
