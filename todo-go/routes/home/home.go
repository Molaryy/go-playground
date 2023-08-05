package home

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getHome(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "Hello World")
}

func GetHomeRoute(router *gin.Engine) {
	router.GET("/", getHome)
}
