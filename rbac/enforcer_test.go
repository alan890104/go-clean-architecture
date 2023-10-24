package rbac

import (
	"log"
	"path/filepath"
	"testing"

	"github.com/casbin/casbin/v2"
	"github.com/stretchr/testify/assert"
)

func NewTestEnforcer() *casbin.Enforcer {
	// Get the absolute path to the project directory
	// projectDir, err := project_root_directory.GetRootDirectory()
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// Use filepath.Join to safely join paths
	modelPath := filepath.Join("model.conf")
	policyPath := filepath.Join("policy.csv")

	enforcer, err := casbin.NewEnforcer(modelPath, policyPath)
	if err != nil {
		log.Fatalf("error: enforcer: %s", err)
	}
	return enforcer
}

func TestAuthorization(t *testing.T) {
	enforcer := NewTestEnforcer()

	tests := []struct {
		Role     string
		Path     string
		Method   string
		Expected bool
	}{
		// Director
		{Director, "/books/1", "DELETE", true},
		{Director, "/books", "PUT", true},
		{Director, "/books", "GET", true},
		{Director, "/records", "GET", true},

		// Librarian
		{Librarian, "/books/1", "DELETE", false},
		{Librarian, "/books", "PUT", true},
		{Librarian, "/books", "GET", true},
		{Librarian, "/records", "GET", true},

		// Visitor
		{Visitor, "/books/1", "DELETE", false},
		{Visitor, "/books", "PUT", false},
		{Visitor, "/books", "GET", true},
		{Visitor, "/records", "GET", false},

		// Unregistered user
		{"", "/books/1", "DELETE", false},
		{"", "/books", "PUT", false},
		{"", "/books", "GET", true},
		{"", "/records", "GET", false},
	}

	for _, test := range tests {
		allowed, err := enforcer.Enforce(test.Role, test.Path, test.Method)
		if err != nil {
			t.Errorf("error: enforcer: %s", err)
		}
		assert.Equal(t, test.Expected, allowed, "Role: %s, Path: %s, Method: %s", test.Role, test.Path, test.Method)
	}
}
