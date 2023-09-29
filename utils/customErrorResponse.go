package utils

import (
	"github.com/labstack/echo/v4"
)

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func CustomErrorResponse(c echo.Context, httpStatus int, errorMessage string) error {
	return c.JSON(httpStatus, errorResponse{
		Code:    httpStatus,
		Message: errorMessage,
	})
}
