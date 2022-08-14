package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context, db *mongo.Client) {
	var body Auth
	var result Auth
	if err := ctx.BindJSON(&body); err != nil {
		ctx.SecureJSON(http.StatusInternalServerError, gin.H{
			"msg": "Something went wrong",
		})
	}

	filter := bson.D{{"email", body.Email}}
	coll := db.Database("chat").Collection("users")
	err := coll.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.SecureJSON(http.StatusUnauthorized, gin.H{"msg": "Invalid email or password"})
			return
		}

		ctx.SecureJSON(http.StatusInternalServerError, gin.H{
			"msg": "Something went wrong",
		})
		return
	}

	if result.Password != body.Password {
		ctx.SecureJSON(http.StatusUnauthorized, gin.H{"msg": "Invalid email or password"})
		return
	}

	ctx.SecureJSON(http.StatusCreated, body)
}
