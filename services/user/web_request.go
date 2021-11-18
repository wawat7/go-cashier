package user


type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name string `json:"name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Address string `json:"address" binding:"required"`
}

type UpdateUserRequest struct {
	Name string `json:"name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Address string `json:"address" binding:"required"`
}