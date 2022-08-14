package main

import (
	"JuHaNi654/server/internal"
	"JuHaNi654/server/routes"
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	db := internal.NewDatabase()

	go internal.InitDatabase(db)
	defer func() {
		db.Disconnect <- true
	}()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to Gin Server changed")

	})

	v1 := router.Group("/v1")
	{
		v1.POST("/login", func(c *gin.Context) { routes.Login(c, db.Client) })
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen %s\n", err)
		}
	}()

	<-ctx.Done()

	stop()
	log.Println("Shutting down gracefully, press Ctrl+c again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
