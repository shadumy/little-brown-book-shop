package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go-with-compose/app/model"
	"go-with-compose/gorm"
	"go-with-compose/mux"
)

type book struct {
	Cover string `json:"cover"`
	Price int    `json:"price"`
	Title string `json:"title"`
	Id    string `json:"id"`
}

type product struct {
	Books []book `json:"books"`
}

// HomePage is the first Page
func HomePage(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the Little Brown Book Shop!")
}

// GetAllBooks is getting list of all available book
func GetAllBooks(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	books := []model.Book{}
	db.Find(&books)

	bookList := []book{}

	for _, s := range books {
		b := book{
			Cover: s.Cover,
			Price: s.Price,
			Title: s.Title,
			Id:    s.Id,
		}
		bookList = append(bookList, b)
	}

	shop := product{
		Books: bookList,
	}

	respondJSON(w, http.StatusOK, shop)
}

// CreateBook is creating a new book to the list
func CreateBook(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	b := model.Book{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&b); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	if err := db.Save(&b).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, b)
}

// GetBook is getting specified book by id
func GetBook(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	b := getBookOr404(db, vars["id"], w, r)
	if b == nil {
		return
	}
	respondJSON(w, http.StatusOK, b)
}

// UpdateBook is updating book information specified by id
func UpdateBook(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	b := getBookOr404(db, vars["id"], w, r)
	if b == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&b); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&b).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, b)

}

// DeleteBook is deleting book specified by id
func DeleteBook(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	b := getBookOr404(db, vars["id"], w, r)

	if b == nil {
		return
	}
	if err := db.Delete(&b).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

// GetDiscount is for calculate discount
func GetDiscount(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	var result map[string][]book
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&result); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	bookList := result["books"]
	net := calcNetDiscount(bookList)
	res := make(map[string]int)
	res["discount"] = net
	respondJSON(w, http.StatusOK, res)

}

// getBookOr404 gets a book instance if exists, or respond the 404 error otherwise
func getBookOr404(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *model.Book {
	book := model.Book{}

	if err := db.Where("id = ?", id).First(&book).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &book
}
