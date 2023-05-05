package biz

import (
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
)


type User struct {
	Email        string
	Token        string
	Username     string
	Bio          string
	Image        string
	PasswordHash string
}

func hashPassowrd(pwd string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func verifyPassowrd(hashed, input string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(input))
	return err == nil 
}

type UserLogin struct {
	Email    string
	Username string
	Token    string
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByUsername(ctx context.Context, username string) (*User, error)
}

type Profile struct {
	Username     string
	Bio          string
	Image        string
	Following    bool
}
type ProfileRepo interface {

}

type UserUsecase struct {
	ur UserRepo
	pr ProfileRepo

	log *log.Helper
}


func NewUserUsecase(ur UserRepo, pr ProfileRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{ur:ur, pr:pr, log: log.NewHelper("xhs", logger)}
}

func (uc *UserUsecase) Register(ctx context.Context, username, email, password string) (*UserLogin, error) {
	u := &User{
		Username: username,
		Email: email,
		PasswordHash: hashPassowrd(password),
	}
	if err := uc.ur.CreateUser(ctx, u); err != nil {
		return nil, err
	} 

	return &UserLogin{
		Email: email,
		Username: username,
		Token: "xxx",
	}, nil 
}

func (uc *UserUsecase) Login(ctx context.Context, email, password string) (*UserLogin, error) {
	u, err :=  uc.ur.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if ok := verifyPassowrd(u.PasswordHash, password); !ok {
		return nil, errors.New("login failed")
	}

	return &UserLogin{
		Email: email,
		Username: u.Username,
		Token: "xxx",
	}, nil 
}