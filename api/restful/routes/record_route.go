package routes

import (
	"github.com/alan890104/go-clean-arch-demo/api/restful/controller"
	"github.com/alan890104/go-clean-arch-demo/api/restful/middleware"
	"github.com/alan890104/go-clean-arch-demo/bootstrap"
	"github.com/gin-gonic/gin"
)

func RegisterRecordRoutes(engine *gin.Engine, recordController *controller.RecordController, app *bootstrap.Application) {
	r := engine.Group("/api/v1/records")
	r.Use(middleware.AuthMiddleware(app.Env.JWT.AccessTokenSecret))
	r.Use(middleware.CasbinRBACMiddleware(app.Enforcer))
	r.GET("", recordController.GetRecord)
	r.GET("/:userID", recordController.GetRecordByUserID)
}
