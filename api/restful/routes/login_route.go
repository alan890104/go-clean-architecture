package routes

import (
	"github.com/alan890104/go-clean-arch-demo/api/restful/controller"
	"github.com/alan890104/go-clean-arch-demo/bootstrap"
)

func RegisterLoginRoutes(app *bootstrap.Application, loginController *controller.LoginController) {
	app.Engine.POST("/login", loginController.Login)
}
