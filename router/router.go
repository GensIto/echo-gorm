package router

import (
	"echo-gorm/usecase"
	"echo-gorm/utils"
	"echo-gorm/validator"
	"net/http"

	validate "github.com/go-playground/validator"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e.Validator = &validator.CustomValidator{Validator: validate.New()}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/auth/login", usecase.Login)
	e.POST("/auth/register", usecase.Register)

	u := e.Group("/users")
	u.Use(echojwt.WithConfig(utils.JwtConfig))
	u.GET("/:id", usecase.GetUser)
	u.DELETE("/:id", usecase.DeleteUser)

	return e
}
