package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Auth struct {
	Email    string
	Password string
}

func Login(ctx *gin.Context, db *mongo.Client) {
	ctx.SecureJSON(http.StatusCreated, gin.H{"msg": "ok"})
}
