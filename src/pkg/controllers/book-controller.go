package controllers

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/xhfmvls/book-management-system/pkg/utils"
	"github.com/xhfmvls/book-management-system/pkg/models"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Parsing error")
	}
	bookDetails, _ := models.GetBookById(id)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// func CreateBook(w http.ResponseWriter, r *http.Request) {
// 	CreateBook := &models.Book{}
// 	utils.ParseBody(r, CreateBook)
// 	book := CreateBook.CreateBook()
// 	res, _ := json.Marshal(book)
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

func CreateBook(w http.ResponseWriter, r *http.Request){
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b:= CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Parsing error")
	}
	book := models.DeleteBookById(id)
	res, _ := json.Marshal(book)
	w.Header().Set("content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	UpdatedBook := &models.Book{}
	utils.ParseBody(r, UpdatedBook)

	vars := mux.Vars(r)
	bookId := vars["id"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Parsing error")
	}

	bookDetails, db := models.GetBookById(id)
	if UpdatedBook.Name != "" {
		bookDetails.Name = UpdatedBook.Name
	}
	if UpdatedBook.Author != "" {
		bookDetails.Author = UpdatedBook.Author
	}
	if UpdatedBook.Publication != "" {
		bookDetails.Publication = UpdatedBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}