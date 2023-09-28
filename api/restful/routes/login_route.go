package routes

import (
	"github.com/alan890104/go-clean-arch-demo/api/restful/controller"
	"github.com/gin-gonic/gin"
)

func RegisterLoginRoutes(engine *gin.Engine, loginController *controller.LoginController) {
	engine.POST("/login", loginController.Login)
}
