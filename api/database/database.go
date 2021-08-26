package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DBConn *gorm.DB

var dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))

type Author struct {
	gorm.Model
	FirstName string
	LastName  string
}

type Article struct {
	gorm.Model
	Title    string
	Content  string
	Author   Author
	AuthorID uint
}

func Init() {
	var err error
	DBConn, err = gorm.Open("postgres", dsn)
	if err != nil {
		panic("Couldn't connect to database:\n" + dsn)
	}
	fmt.Println("Database connected")
	DBConn.AutoMigrate(&Author{}, &Article{})
	fmt.Println("Database migrated")
}
