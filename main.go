package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	handleRequest()
}

type book struct {
	Cover string
	Price int
	Title string
	Id    string
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the Little Brown Book Shop!")
}

func getBookAll(w http.ResponseWriter, r *http.Request) {
	addBook := book{
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9781/4088/9781408855652.jpg",
		Price: 350,
		Title: "Harry Potter and the Philosopher's Stone (I)",
		Id:    "9781408855652",
	}
	json.NewEncoder(w).Encode(addBook)
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/getAddress", getBookAll)
	http.ListenAndServe(":9567", nil)
}
