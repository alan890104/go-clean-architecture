package main

import (
	"github.com/alan890104/go-clean-arch-demo/api/restful/routes"
	"github.com/alan890104/go-clean-arch-demo/bootstrap"
	"github.com/alan890104/go-clean-arch-demo/domain"
	"github.com/alan890104/go-clean-arch-demo/repository"
	"github.com/alan890104/go-clean-arch-demo/repository/mock"
	"github.com/alan890104/go-clean-arch-demo/usecase"
)

var (
	bookRepo   domain.BookRepository
	userRepo   domain.UserRepository
	recordRepo domain.RecordRepository
)

var (
	bookUsecase   domain.BookUsecase
	recordUsecase domain.RecordUsecase
	loginUsecase  domain.LoginUsecase
	sigupUsecase  domain.SignUpUsecase
)

func main() {
	app := bootstrap.App()

	// Repositories
	if app.UseMock {
		bookRepo = mock.NewMockBookRepository()
		userRepo = mock.NewMockUserRepository()
		recordRepo = mock.NewMockRecordRepository()
	} else {
		bookRepo = repository.NewBookRepository(app.Conn)
		userRepo = repository.NewUserRepository(app.Conn)
		recordRepo = repository.NewRecordRepository(app.Conn)
	}

	// Usecases
	bookUsecase = usecase.NewBookUsecase(bookRepo)
	recordUsecase = usecase.NewRecordUsecase(recordRepo)
	loginUsecase = usecase.NewLoginUsecase(userRepo)
	sigupUsecase = usecase.NewSignupUsecase(userRepo)

	routes.RegisterRoutes(app, bookUsecase, recordUsecase, loginUsecase, sigupUsecase)
	app.Run()
}
