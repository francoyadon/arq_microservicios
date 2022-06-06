package router

import (
	bookController "microservicio/controllers/book"
	"github.com/gin-gonic/gin"
	"fmt"
)

func MapUrls(router *gin.Engine) {
	// Products Mapping
	router.GET("/books/:id", bookController.Get)
	router.POST("/books", bookController.Insert)

	fmt.Println("Finishing mappings configurations")
}
