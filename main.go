package main

import (
	"Book/app/middlewares"
	"Book/app/repositories"
	"Book/app/routes"
	"database/sql"
	"fmt"
	"log"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"
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

func main(){
	const port string = ":8080"
	r := gin.Default()
	r.Use(gin.Recovery(), middlewares.Logger())
	 psqlInfo, err := openDB()
	if err != nil {
		log.Printf("error connecting DB: %v", err)
		return
	}
	defer psqlInfo.Close()
	repositories.SetDB(psqlInfo)
	log.Println("DB connection is successful")
	routes.SetupRoutes(r)
	r.Run(port)
}