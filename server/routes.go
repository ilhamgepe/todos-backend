package server

import (
	"github.com/ilhamgepe/todos-backend/internal/repositories"
	"github.com/ilhamgepe/todos-backend/internal/usecases"
	"github.com/ilhamgepe/todos-backend/server/controllers"
)

func (s *Server) SetupRoutes() {
	userRepository := repositories.NewUserRepository(s.DB)
	authUsecase := usecases.NewAuthUsecase(userRepository)
	
	authController := controllers.NewAuthController(authUsecase)

	v1 := s.Gin.Group("/api")

	v1.POST("/auth/register", authController.Register)
	v1.POST("/auth/login", authController.Login)
}