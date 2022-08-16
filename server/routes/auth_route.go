package routes

import (
	"JuHaNi654/server/models"
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-json"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserToken struct {
	Id       string `json:"_id" bson:"_id"`
	Username string `json:"username"`
}

type TokenPayload struct {
	Id string `json:"id"`
	jwt.StandardClaims
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

	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(body.Password)); err != nil {
		log.Println(err)
		ctx.SecureJSON(http.StatusUnauthorized, gin.H{"msg": "Invalid email or password"})
		return
	}

	var payload UserToken
	err = coll.FindOne(context.TODO(), filter).Decode(&payload)

	claims := TokenPayload{
		payload.Id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 6).Unix(),
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

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

func Register(ctx *gin.Context, db *mongo.Client) {
	var body models.RegisterBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]models.FieldErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = models.FieldErrorMsg{
					Field:   fe.Field(),
					Message: models.GetFieldErrorMsg(fe),
				}
			}

			ctx.AbortWithStatusJSON(http.StatusBadRequest, out)
		}
		return
	}

	user := body.SetNewUser()
	b, _ := json.Marshal(user)
	var doc interface{}
	err := bson.UnmarshalExtJSON(b, false, &doc)
	if err != nil {
		log.Println(err)
	}

	coll := db.Database("chat").Collection("users")
	result, err := coll.InsertOne(context.TODO(), doc)

	if err != nil {
		log.Println("Error while saving new user:", err)
		ctx.SecureJSON(http.StatusInternalServerError, gin.H{"msg": "Something went wrong"})
		return
	}

	log.Println("User created: ", result.InsertedID)
	ctx.SecureJSON(http.StatusCreated, gin.H{"msg": "User is successfully created"})
}
