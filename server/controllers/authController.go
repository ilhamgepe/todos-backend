package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ilhamgepe/todos-backend/helper"
	"github.com/ilhamgepe/todos-backend/internal/models"
	"github.com/ilhamgepe/todos-backend/internal/usecases"
)

type AuthController struct {
	authUsecase *usecases.AuthUsecase
}

func NewAuthController(authUsecase *usecases.AuthUsecase) *AuthController{
	return &AuthController{authUsecase: authUsecase}
}

func (ac *AuthController) Register(c *gin.Context){
	var registerData models.UserRegisterDTO
	if err := c.ShouldBind(&registerData); err != nil {
		errorMsg := helper.GenerateMessage(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": errorMsg})
		return
	}

	err := ac.authUsecase.Register(&registerData)
	if err != nil{
		if strings.Contains(err.Error(), "email already exists") {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}



	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (ac *AuthController) Login(c *gin.Context){
	var loginData models.UserLoginDTO
	if err := c.ShouldBind(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res,err  := ac.authUsecase.Login(loginData.Email, loginData.Password)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": res})
}

func (ac *AuthController) Refresh(c *gin.Context){
	claims, err := helper.ValidateRefreshToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"unauthorized"})
		return
	}

	res, err := ac.authUsecase.Refresh(claims.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"unauthorized"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data":res})
}

