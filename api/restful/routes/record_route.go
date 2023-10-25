package routes

import (
	"github.com/alan890104/go-clean-arch-demo/api/restful/controller"
	"github.com/alan890104/go-clean-arch-demo/api/restful/middleware"
	"github.com/alan890104/go-clean-arch-demo/bootstrap"
)

func RegisterRecordRoutes(app *bootstrap.Application, recordController *controller.RecordController) {
	r := app.Engine.Group("/api/v1/records")
	r.Use(middleware.AuthMiddleware(app.Env.JWT.AccessTokenSecret))
	r.Use(middleware.CasbinRBACMiddleware(app.Enforcer))
	r.GET("", recordController.GetRecord)
	r.GET("/:userID", recordController.GetRecordByUserID)
}
