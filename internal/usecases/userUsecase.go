package usecases

import (
	"errors"

	"github.com/ilhamgepe/todos-backend/internal/models"
	"github.com/ilhamgepe/todos-backend/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	ur repositories.UserRepository
}

func NewUserUsecase(ur repositories.UserRepository)*UserUsecase{
	return &UserUsecase{ur}
}


func (u *UserUsecase) CreateUser(user *models.UserRegisterDTO)error{
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	user.Password = string(hashedPassword)

	err = u.ur.Create(user)
	if err != nil {
		return err
	}
	return nil
}