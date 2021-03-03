package internal

import (
	"gorm.io/gorm"
)

type UserService interface {
	GetUser(id uint) string
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return userService{db: db}
}

func (us userService) GetUser(id uint) string {
	return "to be implemented"
}
