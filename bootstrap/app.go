package bootstrap

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	_ "time/tzdata"

	"github.com/alan890104/go-clean-arch-demo/rbac"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AppOpts func(app *Application)

type Application struct {
	Env      *Env
	Conn     *gorm.DB
	Engine   *gin.Engine
	Enforcer *casbin.Enforcer
	UseMock  bool
}

func WithUseMock(useMock bool) AppOpts {
	return func(app *Application) {
		app.UseMock = useMock
	}
}

func App(opts ...AppOpts) *Application {
	env := NewEnv()
	db := NewDB(env)
	engine := gin.Default()
	enforcer := rbac.NewEnforcer()

	// Set timezone
	tz, err := time.LoadLocation(env.Server.TimeZone)
	if err != nil {
		log.Fatal(err)
	}
	time.Local = tz

	app := &Application{
		Env:      env,
		Conn:     db,
		Engine:   engine,
		Enforcer: enforcer,
	}

	for _, opt := range opts {
		opt(app)
	}

	return app
}

// Run run the application with graceful shutdown
func (app *Application) Run() {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.Env.Server.Port),
		Handler: app.Engine,
	}

	// Create a channel to listen for errors coming from the listener. Use a buffered channel so the goroutine can exit if we don't collect this error.
	serverErrors := make(chan error, 1)

	// Start the server in a goroutine so that it doesn't block.
	go func() {
		log.Printf("Server is running on port %d", app.Env.Server.Port)
		serverErrors <- srv.ListenAndServe()
	}()

	// Create a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Block until we receive our signal.
	select {
	case err := <-serverErrors:
		log.Fatalf("Error starting server: %v", err)

	case <-shutdown:
		log.Println("Starting graceful shutdown...")

		// Create a deadline to wait for.
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		// Doesn't block if no connections, but will otherwise wait
		// until the timeout deadline.
		err := srv.Shutdown(ctx)
		if err != nil {
			// If we encounter an error during shutdown, log it and then try to close the server.
			log.Printf("Could not stop server gracefully: %v", err)
			if closeErr := srv.Close(); closeErr != nil {
				// If the server couldn't close, log it as a fatal error.
				log.Fatalf("Could not close http server: %v", closeErr)
			}
		}
	}
}
