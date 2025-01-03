package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TodoItem struct {
	Id 			int 		`json:"id"`
	Title 		string 	  	`json:"title"`
	Description string		`json:"description"`
	Status 		string 		`json:"status"`
	CreatedAt 	time.Time 	`json:"created_at"`
	UpdatedAt 	time.Time 	`json:"updated_at"`
}

type TodoItemCreation struct {
	Title 		string 	  	`json:"title"`
	Description string		`json:"description"`
	Status 		string 		`json:"status"`
}

func main() {
	


	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/api/v1")
	{
		items := v1.Group("/items") 
		{
			items.POST("", createItems)
			items.GET("", getAllItems)
			items.GET(":id", getItemById)
			items.PATCH(":id", editItem)
			items.DELETE(":id", deleteItem)
		}
	}

	r.Run("localhost:8000")
} 

func createItems(c *gin.Context) {
	var data TodoItemCreation 

	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"data": data,
	})	


}

func getAllItems(c *gin.Context) {
	return
}
func getItemById(c *gin.Context) {
	return
}
func editItem(c *gin.Context) {
	return
}
func deleteItem(c *gin.Context) {
	return
}
