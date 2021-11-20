package user

import (
	"github.com/wawat7/go-cashier/helper"
	"gorm.io/gorm"
	"math"
)

type Repository interface {
	Save(user User) User
	Update(user User) User
	Delete(user User)
	FindById(ID int) User
	FindAll() []User
	FindAllWithPagination(page int, pageSize int) ([]User, helper.Pagination)
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

func (repository *repository) FindAllWithPagination(page int, pageSize int) ([]User, helper.Pagination) {
	var users []User
	var total int64
	repository.db.Find(&users).Count(&total)

	err := repository.db.Scopes(Paginate(page, pageSize)).Find(&users).Error
	helper.PanicIfError(err)

	pagination := helper.Pagination{
		Total:       int(total),
		Count:       len(users),
		PerPage:     pageSize,
		CurrentPage: page,
		TotalPage:   int(math.Ceil(float64(total) / float64(pageSize))) ,
	}

	return users, pagination
}

func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}


