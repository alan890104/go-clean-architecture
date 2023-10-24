package rbac

import (
	"log"
	"path/filepath"

	"github.com/casbin/casbin/v2"
	_ "github.com/go-sql-driver/mysql"
	project_root_directory "github.com/golang-infrastructure/go-project-root-directory"
)

func NewEnforcer() *casbin.Enforcer {
	// Get the absolute path to the project directory
	projectDir, err := project_root_directory.GetRootDirectory()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Use filepath.Join to safely join paths
	modelPath := filepath.Join(projectDir, "rbac", "model.conf")
	policyPath := filepath.Join(projectDir, "rbac", "policy.csv")

	enforcer, err := casbin.NewEnforcer(modelPath, policyPath)
	if err != nil {
		log.Fatalf("error: enforcer: %s", err)
	}
	return enforcer
}
