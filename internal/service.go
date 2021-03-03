package internal

import (
	"context"
	"darshanwj/gorm-test/internal/model"
	"log"

	"gorm.io/gorm"
)

type UserService interface {
	GetUser(ctx context.Context, id uint) model.User
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return userService{db: db}
}

func (us userService) GetUser(ctx context.Context, id uint) model.User {
	var user model.User
	err := us.db.WithContext(ctx).Preload("Comments").First(&user, id).Error
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(user)
	return user
}
