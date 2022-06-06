package book

import (
	model "microservicio/models"

)


func GetById(id string) model.Book {
	var book model.Book
 	book.Name="test model mock"
 	book.Id=id
	return book
}
 

func Insert(book model.Book) model.Book {
  	book.Id="1234AA"
	return book
}
