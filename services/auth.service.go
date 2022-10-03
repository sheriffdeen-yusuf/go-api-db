package services

import (
	"github.com/adewumi0550/pro_work/m/v2/repository"
)

// import "os/user"

type AuthServices interface {
	LoginCrendtial(email string, password string) interface{}
	RegisterCredntial(email string, password string, fullName string) interface{}
	ResetPassword(email string) bool
	FindEmailAddress(email string) bool
}

type authServices struct {
	userRepository repository.UserRepository
}

// func NewAuthService(userRep repository.UserRepository) AuthServices {
// 	return &authServices{
// 		userRepository: userRep,
// 	}
// }

// func (services *authServices) LoginCrendtial(email string, password string) interface{}{
// select * from user where email = email
// res := services.

// }
