package usecases

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
	token,refresh,err := createToken(user)
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

func createToken(user *models.User)(*string,*string,error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour).Unix(),
	})
	refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	tokenString,err  := token.SignedString(([]byte(os.Getenv("JWT_SECRET"))))
	if err != nil {
		return nil,nil,err
	}
	refreshString,err  := refresh.SignedString(([]byte(os.Getenv("JWT_SECRET_REFRESH"))))
	if err != nil {
		return nil,nil,err
	}

	return &tokenString,&refreshString,nil
}
