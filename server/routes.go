package server

import (
	"github.com/ilhamgepe/todos-backend/internal/repositories"
	"github.com/ilhamgepe/todos-backend/internal/usecases"
	"github.com/ilhamgepe/todos-backend/server/controllers"
)

func (s *Server) SetupRoutes() {
	// repositories
	userRepository := repositories.NewUserRepository(s.DB)

	// usecases
	userUsecase := usecases.NewUserUsecase(userRepository)
	authUsecase := usecases.NewAuthUsecase(userRepository,userUsecase)
	
	// controllers
	authController := controllers.NewAuthController(authUsecase)

	// routes
	v1 := s.Gin.Group("/api")

	v1.POST("/auth/register", authController.Register)
	v1.POST("/auth/login", authController.Login)
	v1.GET("/auth/refresh", authController.Refresh)
}