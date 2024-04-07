package repositories

import (
	"database/sql"

	"github.com/ilhamgepe/todos-backend/internal/models"
)

type UserRepository interface {
	Create(user *models.UserRegisterDTO) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindById(id int) (*models.User, error)
	Update(user *models.User) (*models.User, error)
}

type userRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
return &userRepositoryImpl{DB: db}
}

func (r *userRepositoryImpl) Create(user *models.UserRegisterDTO) (*models.User, error) {
	var(
		firstName = "ilhamgepe"
		lastName = "ilhamgepe"
	)
	return &models.User{
		ID: 1,
		FirstName: &firstName,
		LastName: &lastName,
		Email: "ilham@gmail.com",
		Password: "ilhamgepe",		
	},nil
}
func (r *userRepositoryImpl) FindByEmail(email string) (*models.User, error) {
	var(
		firstName = "ilhamgepe"
		lastName = "ilhamgepe"
	)
	return &models.User{
		ID: 1,
		FirstName: &firstName,
		LastName: &lastName,
		Email: "ilham@gmail.com",
		Password: "ilhamgepe",		
	},nil
}
func (r *userRepositoryImpl) FindById(id int) (*models.User, error) {
		var(
		firstName = "ilhamgepe"
		lastName = "ilhamgepe"
	)
	return &models.User{
		ID: 1,
		FirstName: &firstName,
		LastName: &lastName,
		Email: "ilham@gmail.com",
		Password: "ilhamgepe",		
	},nil
}
func (r *userRepositoryImpl) Update(user *models.User) (*models.User, error) {
		var(
		firstName = "ilhamgepe"
		lastName = "ilhamgepe"
	)
	return &models.User{
		ID: 1,
		FirstName: &firstName,
		LastName: &lastName,
		Email: "ilham@gmail.com",
		Password: "ilhamgepe",		
	},nil
}
