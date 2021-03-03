package internal

import (
	"context"
	"darshanwj/gorm-test/internal/model"
	"log"

	"gorm.io/gorm"
)

type UserService interface {
	GetUser(ctx context.Context, id uint) model.User
	GetUsers(ctx context.Context) []model.User
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return userService{db: db}
}

func (us userService) GetUser(ctx context.Context, id uint) model.User {
	var user model.User
	err := us.db.WithContext(ctx).Preload("Comments").Take(&user, id).Error
	if err != nil {
		log.Println(err.Error())
	}
	return user
}

func (us userService) GetUsers(ctx context.Context) []model.User {
	var users []model.User
	err := us.db.WithContext(ctx).Preload("Comments").Find(&users).Error
	if err != nil {
		log.Println(err.Error())
	}
	return users
}
