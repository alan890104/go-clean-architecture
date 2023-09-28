package main

import (
	"path"

	"github.com/alan890104/go-clean-arch-demo/bootstrap"
	"github.com/alan890104/go-clean-arch-demo/domain"
	"gorm.io/gen"
)

func main() {
	app := bootstrap.App()
	g := gen.NewGenerator(gen.Config{
		OutPath: path.Join("domain", "query"),
		Mode:    gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	gormdb := bootstrap.NewMySQLDB(app.Env)
	g.UseDB(gormdb)
	g.ApplyBasic(domain.Book{}, domain.User{})
	g.Execute()
}
