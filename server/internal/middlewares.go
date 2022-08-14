package internal

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Authorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := strings.Split(c.Request.Header["Authorization"][0], " ")[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("AllYourBase"), nil
		})

		if token.Valid {
			log.Println("We cool")
			return
		}

		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				log.Println("That's not even a token")
			} else if ve.Errors&jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet != 0 {
				// Token is either expired or not active yet
				log.Println("Timing is everything")
			} else {
				log.Println("Couldn't handle this token: ", err)
			}
		} else {
			log.Println("Couldn't handle this token: ", err)
		}

	}
}
