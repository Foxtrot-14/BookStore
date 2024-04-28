package main

import (
	"book/controller"
	"book/middlewares"
	"book/service"

	"github.com/gin-gonic/gin"
)
var(
    bookService service.BookService = service.New()
    bookController controller.BookController = controller.New(bookService)
)
func main() {
    server := gin.New()
    server.Use(gin.Recovery(), middlewares.Logger())
    server.GET("/books",func (ctx *gin.Context){
        ctx.JSON(200,bookController.FindAll())
    })
    server.POST("/books",func (ctx *gin.Context){
        ctx.JSON(200,bookController.Save(ctx))
    })
    server.Run(":8080")
}
