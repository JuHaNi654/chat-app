package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReturnServerError(err error, ctx *gin.Context) {
	log.Println("Error occurred while handling route: ", err)
	ctx.SecureJSON(http.StatusInternalServerError, gin.H{"msg": "Something went wront"})
}
