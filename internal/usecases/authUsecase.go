package usecases

import (
	"errors"

	"github.com/ilhamgepe/todos-backend/helper"
	"github.com/ilhamgepe/todos-backend/internal/models"
	"github.com/ilhamgepe/todos-backend/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	ur repositories.UserRepository
	uu *UserUsecase
}

func NewAuthUsecase(ur repositories.UserRepository,uu *UserUsecase) *AuthUsecase{
	return &AuthUsecase{ur: ur,uu:uu}
}

func (a *AuthUsecase) Register(user *models.UserRegisterDTO)error{
	err := a.uu.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthUsecase) Login(email string, password string) (*models.LoginResponse,error){
	// check if user exists
	user,err := a.ur.FindByEmail(email)
	if err != nil {
		return nil,err
	}

	// compare password
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil,errors.New("invalid credentials")
	}

	// generate tokens
	token,refresh,err := helper.CreateToken(user)
	if err != nil {
		return nil,errors.New("failed to generate token")
	}

	return &models.LoginResponse{
		User: models.User{
			ID: user.ID,
			Email: user.Email,
			FirstName: user.FirstName,
			LastName: user.LastName,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
		Tokens: models.Tokens{
			AccessToken:  *token,
			RefreshToken: *refresh,
		},
	},nil
}

func (a *AuthUsecase) Update(user *models.User) error {
	return a.ur.Update(user)
}

func (a *AuthUsecase) UpdatePassword(id int, password string) error {
	return a.ur.UpdatePassword(id, password)
}

func (a *AuthUsecase) Refresh(email string) (*models.LoginResponse,error){
	// check if user exists
	user,err := a.ur.FindByEmail(email)
	if err != nil {
		return nil,err
	}

	// generate tokens
	token,refresh,err := helper.CreateToken(user)
	if err != nil {
		return nil,errors.New("failed to generate token")
	}

	return &models.LoginResponse{
		User: models.User{
			ID: user.ID,
			Email: user.Email,
			FirstName: user.FirstName,
			LastName: user.LastName,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
		Tokens: models.Tokens{
			AccessToken:  *token,
			RefreshToken: *refresh,
		},
	},nil
}



