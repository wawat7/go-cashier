package user

import (
	"errors"
	"github.com/wawat7/go-cashier/helper"
	"time"
)

type Service interface {
	Create(input CreateUserRequest, user User) User
	Update(user User) User
	Delete(ID int)
	FindById(ID int) User
	FindAll() []User
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (service *service) Create(input CreateUserRequest, user User) User {
	user = User{
		Username:    input.Username,
		Password:    input.Password,
		Name:        input.Name,
		PhoneNumber: input.PhoneNumber,
		Address:     input.Address,
		Role:        ROLE_ADMIN,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	user = service.repository.Save(user)
	return user
}

func (service *service) Update(user User) User {
	user = service.repository.Update(user)
	return user
}

func (service *service) Delete(ID int) {
	user := service.FindById(ID)
	service.repository.Delete(user)
	return
}

func (service *service) FindById(ID int) User {
	user := service.repository.FindById(ID)
	if user.ID == 0 {
		helper.PanicIfError(errors.New("user not found"))
	}
	return user
}

func (service *service) FindAll() []User {
	return service.repository.FindAll()
}

