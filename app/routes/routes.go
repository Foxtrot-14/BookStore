package routes

import (
	"Book/app/controller"

	"github.com/gin-gonic/gin"
)
func SetupRoutes(router *gin.Engine){
	router.GET("/book",controller.GetAllBooks)
	router.GET("/book/:id",controller.GetBookById)
	router.POST("/book",controller.AddBook)
	router.PUT("/book/:id",controller.UpdateBook)
	router.DELETE("/book/:id",controller.DeleteBookById)	
}