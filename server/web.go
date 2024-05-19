package server

import (
	"github.com/gin-gonic/gin"
	"github.com/hkbiet/golang-basic-auth/basicAuthentication"
	"net/http"
)

func Serve() {
	router := gin.Default()

	router.GET("/profile", basicAuthentication.BasicAuthMiddleware(), func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "You are authenticated"})
	})

	err := router.Run("0.0.0.0:8000")
	if err != nil {
		panic(err)
	}
}
