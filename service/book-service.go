package service

import (
	"book/entity"
)

type BookService interface {
	Save(entity.Book) entity.Book
	FindAll() []entity.Book
}
type bookService struct{
	books []entity.Book
}
func New() BookService {
	return &bookService{}
}
func (service *bookService) Save(book entity.Book) entity.Book {
	service.books = append(service.books, book)
	return book
}
func (service *bookService) FindAll() []entity.Book {
	return service.books
}