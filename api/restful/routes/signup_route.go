package routes

import (
	"github.com/alan890104/go-clean-arch-demo/api/restful/controller"
	"github.com/gin-gonic/gin"
)

func RegisterSignupRoutes(engine *gin.Engine, signupController *controller.SignupController) {
	engine.POST("/signup", signupController.Signup)
}
