package main

import (
	"github.com/alan890104/go-clean-arch-demo/api/restful/routes"
	"github.com/alan890104/go-clean-arch-demo/bootstrap"
	"github.com/alan890104/go-clean-arch-demo/repository"
	"github.com/alan890104/go-clean-arch-demo/usecase"
)

func main() {
	app := bootstrap.App()

	bookRepo := repository.NewMysqlBookRepository(app.Conn)
	bookUsecase := usecase.NewBookUsecase(bookRepo)

	userRepo := repository.NewUserRepository(app.Conn)
	loginUsecase := usecase.NewLoginUsecase(userRepo)
	sigupUsecase := usecase.NewSignupUsecase(userRepo)

	routes.RegisterRoutes(app, bookUsecase, loginUsecase, sigupUsecase)
	app.Run()
}
