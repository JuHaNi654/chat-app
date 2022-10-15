package routes

import (
	"JuHaNi654/server/internal"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(internal.SetHeaders())
	return router
}

func setupTest(tb testing.TB) func(tb testing.TB) {
	log.Println("Setup test")

	return func(tb testing.TB) {
		log.Println("Teardown test")

		// TODO reset mongodb documents after tests

	}
}

func TestLoginRoute(t *testing.T) {
	dbClient := internal.NewDatabase()
	go internal.InitDatabase(dbClient)
	defer func() {
		dbClient.Disconnect <- true
	}()

	r := setupRouter()
	r.POST("/login", func(c *gin.Context) { Login(c, dbClient.Client) })

	body := Auth{
		Email:    "",
		Password: "",
	}

	jsonValue, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
