package controllers

import (
	"fmt"
	"net/http"
	"time"
	"strings"

	"github.com/CalvinLe08/todo-app/models"
	"github.com/CalvinLe08/todo-app/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(DB *gorm.DB) AuthController {
	return AuthController {
		DB: DB,
	}
}

func (ac *AuthController) Register(c *gin.Context) {
	var registerInfo *models.RegisterInput

	if err := c.ShouldBindJSON(&registerInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "Invalid input",
			"details": fmt.Sprintf("Validation failed: %v", err),
		})
		return
	}
	
	if registerInfo.Password != registerInfo.PasswordConfirm {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"message": "Passwords do not match",
		})
		return
	}
	
	hashedPassword, err := utils.HashPassword(registerInfo.Password)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H {
			"status": "error", 
			"message": err.Error(),
		})
		return
	}

	now := time.Now()

	newUser := &models.User{
		ID: uuid.New(),
		Name: registerInfo.Name,
		Email: registerInfo.Email,
		Age: registerInfo.Age,
		Password: hashedPassword,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := ac.DB.Create(&newUser)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			c.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "User already exists"})
			return
		}
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data": newUser,
	})
}

func (ac *AuthController) SignIn(c *gin.Context) {

}


