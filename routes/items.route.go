package routes

import (
	"github.com/calvinnle/todo-app/controllers"
	"github.com/calvinnle/todo-app/middleware"
	"github.com/gin-gonic/gin"
)

type ItemRouteController struct {
	ItemController controllers.ItemController
}

func NewItemRouteController(ItemController controllers.ItemController) ItemRouteController {
	return ItemRouteController{ItemController}
}

func (rc *ItemRouteController) ItemRoute(rg *gin.RouterGroup) {
	router := rg.Group("items")

	router.POST("", middleware.DeserializeUSer(),rc.ItemController.CreateItems)
}

