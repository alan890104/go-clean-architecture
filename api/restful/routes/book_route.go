package routes

import (
	"github.com/alan890104/go-clean-arch-demo/api/restful/controller"
	"github.com/gin-gonic/gin"
)

func RegisterBookRoutes(engine *gin.Engine, bookController *controller.BookController) {
	r := engine.Group("/api/v1/books")
	r.GET("", bookController.GetBook)
	r.GET("/:bookID", bookController.GetBookByID)
	r.PUT("/:bookID", bookController.UpdateBookByID)
	r.DELETE("/:bookID", bookController.DeleteBookByID)
	r.POST("", bookController.CreateBook)
	r.POST("/borrow/:bookID", bookController.BorrowBook)
	r.POST("/return/:bookID", bookController.ReturnBook)
}
