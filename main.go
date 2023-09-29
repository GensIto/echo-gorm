package main

import (
	"echo-gorm/db"
	"echo-gorm/models"
	"echo-gorm/router"
	"echo-gorm/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	utils.LoadEnv()

	db, err := db.NewDB()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{})

	e := router.NewRouter()
	e.Use(middleware.Logger(), middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders,
			echo.HeaderXCSRFToken,
		},
		AllowCredentials: true,
	}))
	e.Logger.Fatal(e.Start(":1323"))
}
