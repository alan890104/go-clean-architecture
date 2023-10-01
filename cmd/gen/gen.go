package main

import (
	"path"

	"github.com/alan890104/go-clean-arch-demo/domain"
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: path.Join("domain", "query"),
		Mode:    gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	g.ApplyBasic(domain.Book{}, domain.User{})
	g.Execute()
}
