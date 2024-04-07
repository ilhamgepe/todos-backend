package usecases

import (
	"errors"

	"github.com/ilhamgepe/todos-backend/internal/models"
	"github.com/ilhamgepe/todos-backend/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	ur repositories.UserRepository
}

func NewAuthUsecase(ur repositories.UserRepository) *AuthUsecase{
	return &AuthUsecase{ur: ur}
}

func (a *AuthUsecase) Register(user *models.UserRegisterDTO)error{
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	user.Password = string(hashedPassword)

	err = a.ur.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthUsecase) Login(email string, password string) (bool,error){
	user,err := a.ur.FindByEmail(email)
	if err != nil {
		return false,err
	}
	if password != user.Password {
		return false,errors.New("password not match")
	}
	return true,nil
}