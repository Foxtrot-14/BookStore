package controller

import (
	"book/entity"
	"book/service"

	"github.com/gin-gonic/gin"
)

type BookController interface {
	FindAll() []entity.Book
	Save(ctx *gin.Context) entity.Book
}
type controller struct {
	service service.BookService
}
func New(service service.BookService) BookController {
	return &controller{
		service: service,
	}
}
func (c *controller) FindAll() []entity.Book {
	return c.service.FindAll()
}
func (c *controller) Save(ctx *gin.Context) entity.Book {
	var book entity.Book
	ctx.BindJSON(&book)
	c.service.Save(book)
	return book
}
