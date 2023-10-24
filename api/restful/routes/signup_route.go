package routes

import (
	"github.com/alan890104/go-clean-arch-demo/api/restful/controller"
	"github.com/alan890104/go-clean-arch-demo/bootstrap"
)

func RegisterSignupRoutes(app *bootstrap.Application, signupController *controller.SignupController) {
	app.Engine.POST("/signup", signupController.Signup)
}
