package models

type User struct {
	ID       uint   `json:"id" gorm:"primary_key" auto_increment:"true"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email" gorm:"unique"`
	PassWord string `json:"password" validate:"required"`
}
