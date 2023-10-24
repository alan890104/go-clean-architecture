package rbac

import (
	"log"

	"github.com/casbin/casbin/v2"
	_ "github.com/go-sql-driver/mysql"
)

func NewEnforcer() *casbin.Enforcer {
	enforcer, err := casbin.NewEnforcer("rbac/model.conf", "rbac/policy.csv")

	if err != nil {
		log.Fatalf("error: enforcer: %s", err)
	}
	return enforcer
}
