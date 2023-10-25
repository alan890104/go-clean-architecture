package rbac

import (
	"log"
	"path/filepath"

	"github.com/casbin/casbin/v2"
)

func NewEnforcer() *casbin.Enforcer {
	modelPath := filepath.Join("rbac", "model.conf")
	policyPath := filepath.Join("rbac", "policy.csv")

	enforcer, err := casbin.NewEnforcer(modelPath, policyPath)
	if err != nil {
		log.Fatalf("error: enforcer: %s", err)
	}
	return enforcer
}
