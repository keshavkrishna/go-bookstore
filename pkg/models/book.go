package models

import (
	"github.com/jinzhu/gorm"
	"github.com/keshavkrishna/go-bookstore/config"
)

var db *gorm.DB 

type Book struct{
	gorm.model
	Name string `gorm:""json:"name""`
	Author string `json:"authr"`
	Publication string `json:"publication"`
}

func init(){
	config.connect()
	db  = config.GetDB()
	db.AutoMigrate(&Book())
}
func (b *Book) CreateBook() *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks(){
	var Books []Book
	db.find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB){
	var getBook Book
	db:=db.Where("Id=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) Book{
	var book Book
	db.Where("Id=?",Id).Delete(book)
	return book
}
