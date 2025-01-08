package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/CalvinLe08/todo-app/initializers"
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
	var SignInInput *models.SignInInput
	
	// Get Sign In input
	if err := c.ShouldBindJSON(&SignInInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {		
			"status": "fail",
			"message": err.Error(),
		})

		return
	}

	// Check db if user existed
	var user models.User

	result := ac.DB.First(&user, "email = ?", strings.ToLower(SignInInput.Email)) 
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail", 
			"message": "Invalid email",
		})
		return
	}

	if err := utils.VerifyPassword(user.Password, SignInInput.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail", 
			"message": "Invalid Password",
		})
		return
	}
	// Get token config
	config, _ := initializers.LoadConfig(".")

	// Generate tokens and return to users
	access_token, err := utils.CreateToken(config.AccessTokenExpiresIn, user.ID, config.AccessTokenPrivate) 
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "logged in",
		"access_token": access_token,
	})
}

func (ac *AuthController) RefreshToken(c *gin.Context) {

}

func (ac *AuthController) SignOut(c *gin.Context) {

}
