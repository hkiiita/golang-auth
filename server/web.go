package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Serve() {
	router := gin.Default()

	router.GET("/profile", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Hello"})
	})

	err := router.Run("0.0.0.0:8000")
	if err != nil {
		panic(err)
	}
}
