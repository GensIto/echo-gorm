package usecase

import (
	"echo-gorm/db"
	"echo-gorm/models"
	"echo-gorm/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {
	claims := utils.GetClaims(c)
	id := claims.ID

	u := models.User{}
	db, err := db.NewDB()
	if err != nil {
		utils.CustomErrorResponse(c, http.StatusInternalServerError, "db error")
	}
	db.First(&u, id)

	return c.JSON(http.StatusOK, u)
}

func DeleteUser(c echo.Context) error {
	claims := utils.GetClaims(c)
	id := claims.ID

	u := models.User{}
	db, err := db.NewDB()
	if err != nil {
		utils.CustomErrorResponse(c, http.StatusInternalServerError, "db error")
	}
	db.Delete(&u, id)

	return c.JSON(http.StatusOK, "seccessfully deleted")
}
