package models

type User struct {
	ID        int     `json:"id"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

type UserRegisterDTO struct {
	FirstName *string `json:"first_name" binding:"omitempty,min=2,max=25"`
	LastName  *string `json:"last_name" binding:"omitempty,min=2,max=25"`
	Email     string  `json:"email" binding:"required,email"`
	Password  string  `json:"password" binding:"required"`
}

type UserLoginDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=100"`
}