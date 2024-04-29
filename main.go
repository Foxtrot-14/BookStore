package main

import (
	"book/controller"
	"book/middlewares"
	"book/service"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)
var(
    bookService service.BookService = service.New()
    bookController controller.BookController = controller.New(bookService)
)
func openDB() (*sql.DB, error) {
	var (
		host     = "localhost"
		port     = 3306
		user     = "root"
		password = "root"
		dbname   = "BookStore"
	)

	// Construct MySQL connection string
	mysqlInfo := os.Getenv("MYSQL_CONN_STRING")
	if len(mysqlInfo) == 0 {
		mysqlInfo = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbname)
	}

	// Open connection to MySQL database
	db, err := sql.Open("mysql", mysqlInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}


func main() {
    server := gin.New()
    server.Use(gin.Recovery(), middlewares.Logger())
    psqlInfo, err := openDB()
	if err != nil {
		log.Printf("error connecting DB: %v", err)
		return
	}
	log.Println("DB connection is successful")
	defer psqlInfo.Close()
    server.GET("/books",func (ctx *gin.Context){
        ctx.JSON(200,bookController.FindAll())
    })
    server.POST("/books",func (ctx *gin.Context){
        ctx.JSON(200,bookController.Save(ctx))
    })
    server.Run(":8080")
}
