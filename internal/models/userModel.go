package models

import "time"

type User struct {
	ID        int     `json:"id"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Email     string  `json:"email"`
	Password  string  `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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

type LoginResponse struct {
	User   User   `json:"user"`
	Tokens Tokens `json:"tokens"`
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
type JWTClaims struct {
	Email string `json:"email"`
	Sub   int    `json:"sub"`
	Exp   float64  `json:"exp"`
}