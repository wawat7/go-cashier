package user

import (
	"github.com/wawat7/go-cashier/helper"
	"gorm.io/gorm"
)

type Repository interface {
	Save(user User) User
	Update(user User) User
	Delete(user User)
	FindById(ID int) User
	FindAll() []User
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (repository *repository) Save(user User) User {
	err := repository.db.Create(&user).Error
	helper.PanicIfError(err)

	return user
}

func (repository *repository) Update(user User) User {
	err := repository.db.Save(&user).Error
	helper.PanicIfError(err)

	return user
}

func (repository *repository) Delete(user User) {
	err := repository.db.Delete(&user).Error
	helper.PanicIfError(err)

	return
}

func (repository *repository) FindById(ID int) User {
	var user User
	err := repository.db.Where("id = ?", ID).Find(&user).Error
	helper.PanicIfError(err)

	return user
}

func (repository *repository) FindAll() []User {
	var users []User
	err := repository.db.Find(&users).Error
	helper.PanicIfError(err)

	return users
}



