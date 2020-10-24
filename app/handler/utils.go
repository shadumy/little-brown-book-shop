package handler

import (
	jwt "go-with-compose/jwt-go"
	"net/http"
	"strings"
)

func calcDiscountPercentage(bookList []book) int {
	hpID := []string{"9781408855652", "9781408855669", "9781408855676",
		"9781408855683", "9781408855690", "9781408855706", "9781408855713"}
	hpCheck := []string{}

	for _, book := range bookList {
		if stringInSlice(book.Id, hpID) {
			hpCheck = append(hpCheck, book.Id)
		}
	}
	hpCheck = uniqueSlice(hpCheck)
	discountRate := []int{0, 0, 10, 11, 12, 13, 14, 15}
	discount := 0
	if len(hpCheck) < 8 {
		discount = discountRate[len(hpCheck)]
	} else {
		discount = 15
	}
	return discount

}

func calcNetDiscount(bookList []book) int {
	discount := calcDiscountPercentage(bookList)
	hpID := []string{"9781408855652", "9781408855669", "9781408855676",
		"9781408855683", "9781408855690", "9781408855706", "9781408855713"}
	var withDiscount int
	for _, book := range bookList {
		if stringInSlice(book.Id, hpID) {
			withDiscount = withDiscount + book.Price
			hpID = removeEleSlice(book.Id, hpID)
		}
	}
	netDiscount := withDiscount * discount / 100
	return int(netDiscount)
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func uniqueSlice(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func removeEleSlice(item string, list []string) []string {
	for i, other := range list {
		if other == item {
			return append(list[:i], list[i+1:]...)
		}
	}
	return list
}

func jwtAuthenVerification(reqToken string) (int, string) {

	if reqToken == "" {
		return http.StatusUnauthorized, "Request need Authorization with token in header"
	}
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) != 2 && splitToken[0] != "Bearer " {
		return http.StatusUnauthorized, "Request need Authorization with token in format: Bearer {token}"
	}
	reqToken = splitToken[1]

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(reqToken, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return http.StatusUnauthorized, err.Error()
		}
		return http.StatusBadRequest, err.Error()
	}
	if !tkn.Valid {
		return http.StatusUnauthorized, "Invalid Signature"
	}
	return http.StatusOK, ""
}
