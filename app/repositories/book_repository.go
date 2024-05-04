package repositories

import (
	"Book/app/models"
	"database/sql"
)

var db *sql.DB;

func SetDB(database *sql.DB){
	db = database
}
func GetAllBooks()([]models.Book,error){
	rows,err := db.Query("SELECT id, title, author FROM books")
	if err!= nil{
		return nil,err
	}
	defer rows.Close()
	books := make([]models.Book,0)
	for rows.Next(){
		var book models.Book
		err := rows.Scan(&book.ID,&book.Title,&book.Author)
		if err!= nil{
			return nil,err
		}
		books = append(books, book)
	}
	return books,nil
}
func GetBookById(id int)(*models.Book,error){
	var book models.Book
	err := db.QueryRow("SELECT id,title,author FROM books WHERE id=?",id).Scan(&book.ID,&book.Title,&book.Author)
	if err!= nil{
		return nil,err
	}
	return &book, nil
}
func AddBook(book *models.Book) error {
	result,err := db.Exec("INSERT INTO books (title,author) VALUES(?,?)",book.Title,book.Author)
	if err != nil{
		return err
	}
	lastInsertedId,err := result.LastInsertId()
	if err != nil{
		return err
	}
	book.ID = int(lastInsertedId)
	return nil
}
func UpdateBook(book *models.Book) error {
	_,err := db.Exec("UPDATE books SET title=?, author=?  WHERE id=?", book.Title,book.Author,book.ID)
	if err != nil{
		return err
	}
	return nil
}
func DeleteBookById(id int) error {
    _, err := db.Exec("DELETE FROM books WHERE id=?", id)
    if err != nil {
        return err
    }
    return nil
}