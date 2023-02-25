package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go_book_management_system/pkg/models"
	"go_book_management_system/pkg/utils"
	"net/http"
	"strconv"
)

//var newBook models.Book

func GetBook(response http.ResponseWriter, request *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(res)
}

func GetBookById(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(res)
}

func CreateBook(writer http.ResponseWriter, reader *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(reader, CreateBook)
	book := CreateBook.CreateBook()
	res, _ := json.Marshal(book)
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
}

func DeleteBook(writer http.ResponseWriter, reader *http.Request) {
	vars := mux.Vars(reader)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	writer.Header().Set("Content-Type", "Application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
}

// UpdateBook make this working with ORM (dont use this approach of
// deleting book with this bookId and creating a new book)
func UpdateBook(writer http.ResponseWriter, reader *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(reader, updateBook)
	vars := mux.Vars(reader)
	bookId := vars["bookId"]
	fmt.Printf("before parsing book id is %s\n", bookId)
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	fmt.Printf("after parsing book id is %d\n", ID)
	bookAfterUpdate := models.UpdateBook(ID, *updateBook)
	res, _ := json.Marshal(bookAfterUpdate)
	writer.Header().Set("Content-Type", "Application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
}
