package bootstrap

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AppOpts func(app *Application)

type Application struct {
	Env     *Env
	Conn    *gorm.DB
	Engine  *gin.Engine
	UseMock bool
}

func WithUseMock(useMock bool) AppOpts {
	return func(app *Application) {
		app.UseMock = useMock
	}
}

func App(opts ...AppOpts) *Application {
	env := NewEnv()
	db := NewMySQLDB(env)
	engine := gin.Default()

	// Set timezone
	tz, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		log.Fatal(err)
	}
	time.Local = tz

	app := &Application{
		Env:    env,
		Conn:   db,
		Engine: engine,
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
	ch := make(chan os.Signal, 1)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	log.Println("Gracefully shutting down...")
}
