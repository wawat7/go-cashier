package app

import (
	"github.com/wawat7/go-cashier/helper"
	"github.com/wawat7/go-cashier/services/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dsn := "root:root@tcp(localhost:3306)/go_cashier?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)

	db.AutoMigrate(&user.User{})
	return db
}
