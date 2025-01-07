package routes

import (
	"github.com/CalvinLe08/todo-app/controllers"
	"github.com/gin-gonic/gin"
)

type AuthRouteController struct {
	authController controllers.AuthController
}

func NewAuthRouteController(authController controllers.AuthController) AuthRouteController {
	return AuthRouteController{authController}
}

func (rc *AuthRouteController) AuthRoute(rg *gin.RouterGroup) {
	router := rg.Group("auth")

	router.POST("register", rc.authController.Register)
	router.POST("signin", rc.authController.SignIn)
	router.POST("refresh-token", rc.authController.RefreshToken)
	router.POST("sign-out", rc.authController.SignOut)
}
