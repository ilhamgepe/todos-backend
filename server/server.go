package server

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ilhamgepe/todos-backend/database"
)

type Server struct {
	*sql.DB
	*validator.Validate
	Gin *gin.Engine
}

func NewServer() *http.Server {
	gin.ForceConsoleColor() //enable color in console
	s := &Server{
		DB: database.DBConnect(),
		Validate: validator.New(),
		Gin: gin.Default(),
	}


	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler: s.Gin,
	}
	s.SetupRoutes()

	// handle jika plikasi di close
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	
	go func(){
		<-c
		log.Println("menerima sinyal SIGINT atau SIGTERM (ctrl+c). shutting down...")

		if err:= s.DB.Close(); err != nil {
			log.Fatal("failed to close DB:", err)
		}
		fmt.Println("DB connection closed.")

		if err := httpServer.Shutdown(context.Background()); err != nil {
			log.Fatal("failed to shutdown Server:", err)
		}
	}()




	return httpServer
}