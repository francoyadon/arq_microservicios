package main

import (
	"microservicio/router"
	"github.com/gin-gonic/gin"
	"fmt"
)

var (
	gin_router *gin.Engine
)

func main() {
	gin_router = gin.Default()
	router.MapUrls(gin_router)

	fmt.Println("Starting server")
	gin_router.Run(":8090")
}

 