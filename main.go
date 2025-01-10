package main

import (
	"log"
	"net/http"

	"github.com/calvinnle/todo-app/controllers"
	"github.com/calvinnle/todo-app/initializers"
	"github.com/calvinnle/todo-app/routes"
	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine

	AuthController controllers.AuthController
	AuthRouteController routes.AuthRouteController

	ItemController controllers.ItemController
	ItemRouteController routes.ItemRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")

	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	AuthController = controllers.NewAuthController(initializers.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController) 

	ItemController = controllers.NewItemController(initializers.DB)
	ItemRouteController = routes.NewItemRouteController(ItemController)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}
	
	router := server.Group("/api")
	router.GET("/healthcheck", func(c *gin.Context) {
		message := "Welcome to my todo app which is done by AI"
		c.JSON(http.StatusOK, gin.H{
			"message": message,
		})
	})

	AuthRouteController.AuthRoute(router)
	ItemRouteController.ItemRoute(router)
	
	log.Fatal(server.Run("localhost:" + config.ServerPort))
}
