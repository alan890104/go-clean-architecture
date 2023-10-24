package rbac

import (
	"log"
	"path/filepath"

	"github.com/casbin/casbin/v2"
	projectRootDir "github.com/golang-infrastructure/go-project-root-directory"
)

func NewEnforcer() *casbin.Enforcer {
	// Get the absolute path to the project directory
	projectDir, err := projectRootDir.GetRootDirectory()
	if err != nil {
		log.Fatal(err.Error())
	}

	modelPath := filepath.Join(projectDir, "rbac", "model.conf")
	policyPath := filepath.Join(projectDir, "rbac", "policy.csv")

	enforcer, err := casbin.NewEnforcer(modelPath, policyPath)
	if err != nil {
		log.Fatalf("error: enforcer: %s", err)
	}
	return enforcer
}
