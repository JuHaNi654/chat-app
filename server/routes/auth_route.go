package routes

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserToken struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

func Login(ctx *gin.Context, db *mongo.Client) {
	var body Auth
	var result Auth
	var err error
	if err := ctx.BindJSON(&body); err != nil {
		ctx.SecureJSON(http.StatusInternalServerError, gin.H{
			"msg": "Something went wrong",
		})
	}

	filter := bson.D{{"email", body.Email}}
	coll := db.Database("chat").Collection("users")
	err = coll.FindOne(context.TODO(), filter).Decode(&result)

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

	var payload UserToken
	err = coll.FindOne(context.TODO(), filter).Decode(&payload)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  payload.Id,
		"exp": 21600,
	})

	secret := []byte(os.Getenv("JWT_SECRET"))
	tokenString, err := token.SignedString(secret)
	if err != nil {
		log.Printf("Error occurred while creating jwt: %v\n", err)
		ctx.SecureJSON(http.StatusInternalServerError, gin.H{
			"msg": "Something went wrong",
		})
		return
	}

	ctx.SecureJSON(http.StatusCreated, gin.H{"token": tokenString, "username": payload.Username})
}
