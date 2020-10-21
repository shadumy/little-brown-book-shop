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

func TestRemoveEleSlice(t *testing.T) {
	a := []string{"apple", "melon", "banana"}
	aCut := removeEleSlice("apple", a)
	if len(aCut) != 2 {
		t.Error("a length is not matched", aCut)
	}
	for _, ele := range aCut {
		if ele == "apple" {
			t.Error("Still have apple left in a", aCut)
		}
	}
}

func TestDiscountShouldReturn0(t *testing.T) {
	bookList := []book{}

	d := calcDiscountPercentage(bookList)
	if d != 0 {
		t.Error("discount for non-harry-potter book should be 0")
	}

	n := calcNetDiscount(bookList)
	if n != 0 {
		t.Error("net for non-harry-potter book should be 0 but found", n)
	}

}

func TestDiscountHP3(t *testing.T) {
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

	d := calcDiscountPercentage(bookList)
	if d != 11 {
		t.Error("discount for 3 unique harry-potter books should be 11")
	}

	n := calcNetDiscount(bookList)
	if n != 123 {
		t.Error("net for 3 unique harry-potter books should be 123 but found", n)
	}

}

func TestDiscountHP3WithOther(t *testing.T) {
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

	bookList = append(bookList, book{
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9781/4732/9781473225046.jpg",
		Price: 320,
		Title: "Mistborn: Secret History",
		Id:    "9781473225046",
	})

	d := calcDiscountPercentage(bookList)
	if d != 11 {
		t.Error("discount for 3 unique harry-potter books should be 11")
	}

	n := calcNetDiscount(bookList)
	if n != 123 {
		t.Error("net for 3 unique harry-potter books should be 123 but found", n)
	}

}

func TestDiscountHP3Duplicate(t *testing.T) {
	bookList := []book{}
	bookList = append(bookList, book{
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9781/4088/9781408855683.jpg",
		Price: 360,
		Title: "Harry Potter and the Goblet of Fire (IV)",
		Id:    "9781408855683",
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
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9781/4732/9781473225046.jpg",
		Price: 320,
		Title: "Mistborn: Secret History",
		Id:    "9781473225046",
	})

	d := calcDiscountPercentage(bookList)
	if d != 11 {
		t.Error("discount for 3 unique harry-potter books should be 11")
	}

	n := calcNetDiscount(bookList)
	if n != 123 {
		t.Error("net for 3 unique harry-potter books should be 123 but found", n)
	}

}

func TestDiscountHP7(t *testing.T) {
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

	d := calcDiscountPercentage(bookList)
	if d != 15 {
		t.Error("discount for 7 unique harry-potter books should be 15", d)
	}

	n := calcNetDiscount(bookList)
	if n != 384 {
		t.Error("net for 7 unique harry-potter books should be 384 but found", n)
	}

}
