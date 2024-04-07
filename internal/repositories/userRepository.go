package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/ilhamgepe/todos-backend/internal/models"
)

type UserRepository interface {
	Create(user *models.UserRegisterDTO)error
	FindByEmail(email string) (*models.User, error)
	FindById(id int) (*models.User, error)
	Update(user *models.User)  error
	UpdatePassword(id int, password string) error
}

type userRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
return &userRepositoryImpl{DB: db}
}

func (r *userRepositoryImpl) Create(user *models.UserRegisterDTO) error {
	sql := "INSERT INTO users (first_name, last_name, email, password) VALUES (?,?,?,?)"
	_, err := r.DB.Exec(sql, user.FirstName, user.LastName, user.Email, user.Password)

	if err != nil {
		// Mengecek apakah kesalahan disebabkan oleh duplikat entri
		if strings.Contains(err.Error(), "Duplicate entry") {
			return errors.New("email already exists")
		}
		return errors.New(fmt.Sprintf("failed to create user: %v", err.Error()))
	}

	return nil
}

func (r *userRepositoryImpl) FindByEmail(email string) (*models.User, error) {
	var user models.User
	q := "SELECT * FROM users WHERE email = ?"
	if err := r.DB.QueryRow(q, email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) FindById(id int) (*models.User, error) {
	var user models.User
	q := "SELECT * FROM users WHERE id = ?"
	if err := r.DB.QueryRow(q, id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}
func (r *userRepositoryImpl) Update(user *models.User) error {
	existingUser,err := r.FindById(user.ID)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return errors.New(fmt.Sprintf("User with ID %d not found", user.ID))
	}


	// Update the user in the database
	q := "UPDATE users SET first_name = ?, last_name = ?, email = ? WHERE id = ?"
	_,err = r.DB.Exec(q, user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepositoryImpl) UpdatePassword(id int, password string) error {
	existingUser,err := r.FindById(id)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return errors.New(fmt.Sprintf("User with ID %d not found", id))
	}
	// update password
	q := "UPDATE users SET password = ? where id = ?"
	_,err = r.DB.Exec(q, password, id)
	if err != nil {
		return err
	}

	return nil
}