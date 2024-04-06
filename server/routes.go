package server

import "github.com/gin-gonic/gin"

func (s *Server) SetupRoutes() {
	s.Gin.GET("/", func(ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
}