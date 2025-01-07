package controllers

import (
	"fmt"
	"net/http"
	"time"
	"strings"

	"github.com/CalvinLe08/todo-app/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ItemController struct {
	DB *gorm.DB
}

func NewItemController(DB *gorm.DB) ItemController {
	return ItemController {
		DB: DB,
	}
}

func (ic *ItemController) CreateItems(c *gin.Context) {
	// currentUser := c.MustGet("currentUser").(models.User)

	var payload *models.ItemCreation 

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
			"details": fmt.Sprintf("Validation failed: %v", err),
		})
		return
	}

	now := time.Now();

	newItem := &models.Item{
		ID: uuid.New(),
		Title: payload.Title,
		Description: payload.Description,
		// UserID: currentUser.ID,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := ic.DB.Create(&newItem)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			c.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Post with that title already exists"})
			return
		}
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "sucess",
		"data": newItem,
	})
}
