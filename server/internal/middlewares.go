package internal

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Authorized() gin.HandlerFunc {
	secret := []byte(os.Getenv("JWT_SECRET"))

	return func(c *gin.Context) {
		bearer := c.Request.Header["Authorization"]

		if len(bearer) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg": "Unauthorized",
			})
			return
		}

		tokenString := strings.Split(bearer[0], " ")

		if len(tokenString) < 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg": "Unauthorized",
			})
			return
		}

		token, err := jwt.Parse(tokenString[1], func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})

		if token.Valid {
			c.Next()
			return
		} else {
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
						"msg": "Session expired",
					})

					return
				}
			}

			log.Println("Couldn't handle this token: ", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"msg": "Something went wrong",
			})
		}

	}
}
