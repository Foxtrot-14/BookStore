package controller

import (
	"Book/app/models"
	"Book/app/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllBooks(c *gin.Context){
	books,err := services.GetAllBooks()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200,books)
}
func GetBookById(c *gin.Context){
	bookID := c.Param("id")
	id,err := strconv.Atoi(bookID)
	if err!= nil {
		c.JSON(400,gin.H{"error":"Invalid book ID"})
		return
	}
	book,err := services.GetBookById(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200,book)
}
func AddBook(c *gin.Context){
	var book models.Book
	err := c.ShouldBindJSON(&book)
	if err!= nil {
		c.JSON(400,gin.H{"error":err.Error()})
		return
	}
	err = services.AddBook(&book)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201,gin.H{"message":"Book added sucessfully","book":book})	
}
func UpdateBook(c *gin.Context){
	bookId := c.Param("id")
	id,err := strconv.Atoi(bookId)
	if err!= nil {
		c.JSON(400,gin.H{"error":"Invalid Book Id"})
		return
	}
	var updatedBook models.Book
	err = c.ShouldBindJSON(&updatedBook)
	if err != nil{
		c.JSON(400,gin.H{"error":err.Error()})
		return
	}
	updatedBook.ID = id
	err = services.UpdateBook(&updatedBook)
	if err != nil{
		c.JSON(500, gin.H{"error": err.Error()})
        return
	}
	c.JSON(200, gin.H{"message": "Book updated successfully", "book": updatedBook})
}
func DeleteBookById(c *gin.Context){
	bookID := c.Param("id")
	id, err := strconv.Atoi(bookID)
	if err != nil {
        c.JSON(400, gin.H{"error": "Invalid book ID"})
        return
    }
	err = services.DeleteBookById(id)
	if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
	c.JSON(200, gin.H{"message": "Book deleted successfully", "bookID": id})
}