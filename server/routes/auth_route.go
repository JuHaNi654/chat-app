package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Auth struct {
	Email    string
	Password string
}

func Login(ctx *gin.Context, db *mongo.Client) {
	fmt.Println("Login route called")
	var result Auth
	filter := bson.D{{"email", "admin@admin.com"}}
	coll := db.Database("chat").Collection("users")
	err := coll.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		log.Printf("Something went wrong \n")
		log.Println(err)
	}

	ctx.SecureJSON(http.StatusCreated, result)
}
