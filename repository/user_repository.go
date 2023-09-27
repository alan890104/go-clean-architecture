package repository

import "github.com/alan890104/go-clean-arch-demo/domain/query"

type userRepository struct {
	query *query.Query
}
