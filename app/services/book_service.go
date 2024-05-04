package services

import (
	"Book/app/models"
	"Book/app/repositories"
)
func GetAllBooks()([]models.Book,error){
	return repositories.GetAllBooks()
}
func GetBookById(id int)(*models.Book,error){
	return repositories.GetBookById(id)
}
func AddBook(book *models.Book) error {
	return repositories.AddBook(book)
}
func UpdateBook(book *models.Book) error {
	return repositories.UpdateBook(book)
}
func DeleteBookById(id int) error {
    return repositories.DeleteBookById(id)
}