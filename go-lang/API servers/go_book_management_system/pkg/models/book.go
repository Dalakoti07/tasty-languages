package models

import (
	"go_book_management_system/pkg/config"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	// AutoMigrate run auto migration for given models, will only add missing fields, won't delete/change current data
	db.AutoMigrate(&Book{})
}

func (book *Book) CreateBook() *Book {
	// Create insert the value into database
	db.Create(&book)
	return book
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

// DeleteBook would do soft delete
func DeleteBook(ID int64) Book {
	var book Book
	db.Clauses(clause.Returning{}).Where("ID=?", ID).Delete(&book)
	return book
}

func UpdateBook(ID int64, updatedBook Book) Book {
	bookDetails, _ := GetBookById(ID)

	db.Model(bookDetails).Where("id", bookDetails.ID).Updates(
		Book{
			Author:      updatedBook.Author,
			Name:        updatedBook.Name,
			Publication: updatedBook.Publication,
		})

	return *bookDetails
}
