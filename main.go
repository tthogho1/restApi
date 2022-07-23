package main

import (
	"restApi/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/Api", controller.Getting)
	router.Run()
}
