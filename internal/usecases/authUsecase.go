package usecases

import (
	"errors"

	"github.com/ilhamgepe/todos-backend/internal/models"
	"github.com/ilhamgepe/todos-backend/internal/repositories"
)

type AuthUsecase struct {
	ur repositories.UserRepository
}

func NewAuthUsecase(ur repositories.UserRepository) *AuthUsecase{
	return &AuthUsecase{ur: ur}
}

func (a *AuthUsecase) Register(user *models.UserRegisterDTO) (*models.User,error){
	createdUser,err := a.ur.Create(user)
	if err != nil {
		return nil,err
	}
	return createdUser,nil
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