package data

import (
	"context"

	"kratos-bbs-demo/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)


type User struct {
	gorm.Model

	Email string
	Username string
	Bio string
	Image string
}

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) error {
	user := User{
		Email: u.Email,
	}
	return r.data.db.WithContext(ctx).Create(&user).Error
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (*biz.User, error) {
	user := new(biz.User)
	if err := r.data.db.WithContext(ctx).Where("email = ?", email).Find(user).Error; err != nil {
		return nil, err
	}

	return user, nil 
}

func (r *userRepo) GetUserByUsername(ctx context.Context, username string) (*biz.User, error) {
	return nil, nil 
}

type profileRepo struct {
	data *Data
	log  *log.Helper
}

func NewProfileRepo(data *Data, logger log.Logger) biz.ProfileRepo {
	return &profileRepo{
		data: data,
		log:  log.NewHelper("xhs", logger),
	}
}
