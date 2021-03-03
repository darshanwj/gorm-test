package internal

import (
	"context"
	"darshanwj/gorm-test/internal/model"
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserService interface {
	GetUser(ctx context.Context, id uint) model.User
	GetUsers(ctx context.Context) []model.User
	CreateUser(ctx context.Context, cur createUserRequest) model.User
}

type userService struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewUserService(db *gorm.DB, log *zap.Logger) UserService {
	return userService{db: db, log: log}
}

func (us userService) GetUser(ctx context.Context, id uint) model.User {
	var user model.User
	err := us.db.WithContext(ctx).Preload("Comments").Take(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			us.log.Warn("could not find user in db", zap.Uint("id", id))
		} else {
			us.log.Error("db error getting user with comments", zap.Error(err))
		}
	}
	return user
}

func (us userService) GetUsers(ctx context.Context) []model.User {
	var users []model.User
	err := us.db.WithContext(ctx).Preload("Comments").Find(&users).Error
	if err != nil {
		us.log.Error("db error getting users with comments", zap.Error(err))
	}
	return users
}

func (us userService) CreateUser(ctx context.Context, cur createUserRequest) model.User {
	user := model.User{
		Name:  cur.Name,
		Phone: cur.Phone,
	}
	res := us.db.WithContext(ctx).Create(&user)
	if res.Error != nil {
		us.log.Error("db error when creating user", zap.Error(res.Error))
	}
	return user
}
