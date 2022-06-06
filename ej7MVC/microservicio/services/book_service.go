package services

import (
	bookDao "microservicio/daos/book"
	"microservicio/dtos"
	model "microservicio/models"
	e "microservicio/utils/errors"
)

type bookService struct{}

type bookServiceInterface interface {
	GetBook(id string) (dtos.BookDto, e.ApiError)
	 
	InsertBook(bookDto dtos.BookDto) (dtos.BookDto, e.ApiError)
}

var (
	BookService bookServiceInterface
)

func init() {
	BookService = &bookService{}
}

func (s *bookService) GetBook(id string) (dtos.BookDto, e.ApiError) {

	var book model.Book = bookDao.GetById(id)
	var bookDto dtos.BookDto

	if book.Id == "" {
		return bookDto, e.NewBadRequestApiError("book not found")
	}
	bookDto.Name = book.Name
	bookDto.Id = book.Id
	return bookDto, nil
}

 

func (s *bookService) InsertBook(bookDto dtos.BookDto) (dtos.BookDto, e.ApiError) {

	var book model.Book

	book.Name = bookDto.Name 

	book = bookDao.Insert(book)

	bookDto.Id = book.Id

	return bookDto, nil
}
