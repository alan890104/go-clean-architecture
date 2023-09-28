package routes

import (
	"github.com/alan890104/go-clean-arch-demo/api/restful/controller"
	"github.com/gin-gonic/gin"
)

func RegisterBookRoutes(engine *gin.Engine, bookController *controller.BookController) {
	r := engine.Group("/api/v1/books")
	r.GET("", bookController.GetBook)
	r.GET("/:bookID", bookController.GetBookByID)
	r.POST("", bookController.CreateBook)
	r.PUT("/borrow/:bookID", bookController.BorrowBook)
	r.PUT("/return/:bookID", bookController.ReturnBook)
}
