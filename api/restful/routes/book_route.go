package routes

import (
	"github.com/alan890104/go-clean-arch-demo/api/restful/controller"
	"github.com/alan890104/go-clean-arch-demo/api/restful/middleware"
	"github.com/alan890104/go-clean-arch-demo/bootstrap"
)

func RegisterBookRoutes(app *bootstrap.Application, bookController *controller.BookController) {
	r := app.Engine.Group("/api/v1/books")
	r.Use(middleware.AuthMiddleware(app.Env.JWT.AccessTokenSecret))
	r.Use(middleware.CasbinRBACMiddleware(app.Enforcer))
	r.GET("", bookController.GetBook)
	r.GET("/:bookID", bookController.GetBookByID)
	r.PUT("/:bookID", bookController.UpdateBookByID)
	r.DELETE("/:bookID", bookController.DeleteBookByID)
	r.POST("", bookController.CreateBook)
	r.POST("/borrow", bookController.BorrowBook)
	r.POST("/return", bookController.ReturnBook)
}
