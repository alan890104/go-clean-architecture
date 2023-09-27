package routes

import (
	"github.com/alan890104/go-clean-arch-demo/api/restful/controller"
	"github.com/alan890104/go-clean-arch-demo/api/restful/middleware"
	"github.com/alan890104/go-clean-arch-demo/domain"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine, bookUsecase domain.BookUsecase) {
	// Register Global Middleware
	cors := middleware.CORSMiddleware()
	engine.Use(cors)

	// Register Book Routes
	bookController := controller.NewBookController(bookUsecase)
	registerBookRoutes(engine, bookController)
}
