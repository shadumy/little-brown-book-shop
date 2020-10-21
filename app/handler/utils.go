package handler

func calcDiscount(bookList []book) int {
	hpID := []string{"9781408855652", "9781408855669", "9781408855676",
		"9781408855683", "9781408855690", "9781408855706", "9781408855713"}
	hpCheck := []string{}
	for _, book := range bookList {
		if stringInSlice(book.Id, hpID) {
			hpCheck = append(hpCheck, book.Id)
		}
	}
	hpCheck = uniqueSlice(hpCheck)
	if len(hpCheck) > 6 {
		return 15
	}
	return 0

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
