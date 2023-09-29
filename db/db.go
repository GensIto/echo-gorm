package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	// SQLiteデータベースに接続
	db, err := gorm.Open(sqlite.Open("go.db"), &gorm.Config{})
	if err != nil {
		return db, err
	}
	return db, nil
}
