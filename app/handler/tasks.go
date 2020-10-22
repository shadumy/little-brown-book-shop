package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go-with-compose/app/model"
	"go-with-compose/gorm"
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

// GetAllBooks is get list of all available book
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
