package book

import (
	"github.com/gin-gonic/gin"
	"net/http"	
	service "microservicio/services"
	 "microservicio/dtos"
	 "fmt"
)

func Get(c *gin.Context) {

	id:=c.Param("id")
	bookDto, er := service.BookService.GetBook(id)


	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusOK,bookDto)
}

func Insert(c *gin.Context) {
	var bookDto dtos.BookDto
	err := c.BindJSON(&bookDto)

	// Error Parsing json param
	if err != nil {
		
    	fmt.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	bookDto, er := service.BookService.InsertBook(bookDto)

	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, bookDto)
}