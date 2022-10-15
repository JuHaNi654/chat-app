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

func SetupRouter(db *internal.Database) *gin.Engine {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to Gin Server changed")

	})

	router.Use(internal.SetHeaders())

	v1 := router.Group("/v1")
	{
		v1.POST("/login", func(c *gin.Context) { routes.Login(c, db.Client) })
		v1.POST("/register", func(c *gin.Context) { routes.Register(c, db.Client) })
		authorized := v1.Group("/")
		authorized.Use(internal.Authorized())
		authorized.GET("/test", func(c *gin.Context) {
			c.SecureJSON(http.StatusOK, gin.H{"msg": "Authorized route"})
		})

	}

	return router
}

func main() {
	client := internal.NewDatabase()

	go internal.InitDatabase(client)
	defer func() {
		client.Disconnect <- true
	}()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	router := SetupRouter(client)
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
