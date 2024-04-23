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
func (service *bookService) Save(entity.Book) entity.Book {

}
func (service *bookService) FindAll(entity.Book) entity.Book {
	
}