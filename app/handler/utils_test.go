package handler

import (
	"testing"
)

func TestStringInSlice(t *testing.T) {
	a := []string{"1", "2"}

	if !stringInSlice("1", a) {
		t.Error("1 is in a")
	}
	if stringInSlice("0", a) {
		t.Error("0 is not in a")
	}
	if stringInSlice("01", a) {
		t.Error("01 is not in a")
	}
}

func TestUniqueSlice(t *testing.T) {
	a := []string{"1", "2", "1", "3", "2"}
	uniqA := uniqueSlice(a)
	if len(uniqA) != 3 {
		t.Error("a is not unique", uniqA)
	}
}

func TestDiscountShouldReturn0(t *testing.T) {
	bookList := []book{}
	bookList = append(bookList, book{
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9781/4088/9781408855683.jpg",
		Price: 360,
		Title: "Harry Potter and the Goblet of Fire (IV)",
		Id:    "9781408855683",
	})

	bookList = append(bookList, book{
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9781/4088/9781408855690.jpg",
		Price: 380,
		Title: "Harry Potter and the Order of the Phoenix (V)",
		Id:    "9781408855690",
	})

	bookList = append(bookList, book{
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9781/4088/9781408855706.jpg",
		Price: 380,
		Title: "Harry Potter and the Half-Blood Prince (VI)",
		Id:    "9781408855706",
	})

	d := calcDiscount(bookList)
	if d != 0 {
		t.Error("discount for non-harry-potter books lesss than 7 should be 0")
	}

}

func TestDiscountShouldReturn15(t *testing.T) {
	bookList := []book{}
	bookList = append(bookList, book{
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9781/4088/9781408855652.jpg",
		Price: 350,
		Title: "Harry Potter and the Philosopher's Stone (I)",
		Id:    "9781408855652",
	})
	bookList = append(bookList, book{
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9781/4088/9781408855669.jpg",
		Price: 350,
		Title: "Harry Potter and the Chamber of Secrets (II)",
		Id:    "9781408855669",
	})
	bookList = append(bookList, book{
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9781/4088/9781408855676.jpg",
		Price: 340,
		Title: "Harry Potter and the Prisoner of Azkaban (III)",
		Id:    "9781408855676",
	})
	bookList = append(bookList, book{
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9781/4088/9781408855683.jpg",
		Price: 360,
		Title: "Harry Potter and the Goblet of Fire (IV)",
		Id:    "9781408855683",
	})

	bookList = append(bookList, book{
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9781/4088/9781408855690.jpg",
		Price: 380,
		Title: "Harry Potter and the Order of the Phoenix (V)",
		Id:    "9781408855690",
	})

	bookList = append(bookList, book{
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9781/4088/9781408855706.jpg",
		Price: 380,
		Title: "Harry Potter and the Half-Blood Prince (VI)",
		Id:    "9781408855706",
	})

	bookList = append(bookList, book{
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9781/4088/9781408855713.jpg",
		Price: 400,
		Title: "Harry Potter and the Deathly Hallows (VII)",
		Id:    "9781408855713",
	})

	d := calcDiscount(bookList)
	if d == 0 {
		t.Error("discount for non-harry-potter books lesss than 7 should be 0")
	}

}
