package routes

import (
	"github.com/alan890104/go-clean-arch-demo/api/restful/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRecordRoutes(engine *gin.Engine, recordController *controller.RecordController) {
	r := engine.Group("/api/v1/records")
	r.GET("", recordController.GetRecord)
	r.GET("/:userID", recordController.GetRecordByUserID)
}
