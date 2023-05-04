package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)


type User struct {

}
type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
}

type Profile struct {

}
type ProfileRepo interface {
	
}

type UserUsecase struct {
	ur UserRepo
	pr ProfileRepo

	log *log.Helper
}


func NewUserUsecase(ur UserRepo, pr ProfileRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{ur:ur, pr:pr, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) Register(ctx context.Context, u *User) error {
	if err := uc.ur.CreateUser(ctx, u); err != nil {
		return nil 
	} 

	return nil
}