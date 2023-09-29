package usecase

import (
	"echo-gorm/db"
	"echo-gorm/models"
	"echo-gorm/utils"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Register(c echo.Context) error {
	u := models.User{}

	if err := c.Bind(&u); err != nil {
		return err
	}

	db, err := db.NewDB()
	if err != nil {
		utils.CustomErrorResponse(c, http.StatusInternalServerError, "db error")
	}

	existingUser := models.User{}
	if err := db.Where("email = ?", u.Email).First(&existingUser).Error; err == nil {
		return utils.CustomErrorResponse(c, http.StatusBadRequest, "email already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return utils.CustomErrorResponse(c, http.StatusInternalServerError, "db error")
	}

	if err := db.Create(&u).Error; err != nil {
		return utils.CustomErrorResponse(c, http.StatusInternalServerError, "db error")
	}

	return c.JSON(http.StatusOK, u)
}
func Login(c echo.Context) error {
	inputUser := models.User{}

	if err := c.Bind(&inputUser); err != nil {
		return utils.CustomErrorResponse(c, http.StatusInternalServerError, "db error")
	}

	db, err := db.NewDB()
	if err != nil {
		return utils.CustomErrorResponse(c, http.StatusInternalServerError, "db error")
	}

	u := models.User{}
	if err := db.Where("email = ?", inputUser.Email).First(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.CustomErrorResponse(c, http.StatusUnauthorized, "invalid email or password")
		}
		return utils.CustomErrorResponse(c, http.StatusInternalServerError, "db error")
	}

	if u.PassWord != inputUser.PassWord {
		return utils.CustomErrorResponse(c, http.StatusUnauthorized, "invalid email or password")
	}

	token := utils.GenerateToken(u.ID)

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
		"user":  u,
	})
}
