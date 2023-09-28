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

type Application struct {
	Env    *Env
	Conn   *gorm.DB
	Engine *gin.Engine
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

func App() *Application {
	env := NewEnv()
	db := NewMySQLDB(env)
	engine := gin.Default()

	// Set timezone
	tz, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		log.Fatal(err)
	}
	time.Local = tz

	return &Application{
		Env:    env,
		Conn:   db,
		Engine: engine,
	}
}
