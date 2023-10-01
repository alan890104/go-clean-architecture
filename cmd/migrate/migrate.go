package main

import (
	"log"

	"github.com/alan890104/go-clean-arch-demo/bootstrap"
	"github.com/alan890104/go-clean-arch-demo/domain"
)

func main() {
	env := bootstrap.NewEnv()
	db := bootstrap.NewDB(env)

	if err := db.AutoMigrate(&domain.Book{}, &domain.User{}); err != nil {
		log.Fatal(err)
	}
}
