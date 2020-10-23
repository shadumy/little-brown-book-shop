package model

import (
	"go-with-compose/gorm"
	_ "go-with-compose/gorm/dialects/mysql"
)

type Book struct {
	Id    string `gorm:"unique;PRIMARY_KEY" json:"id"`
	Cover string `gorm:"default:null" json:"cover"`
	Price int    `json:"price" json:"price"`
	Title string `gorm:"default:null" json:"title"`
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Book{})

	db.Create(&Book{Id: "9781408855652",
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9781/4088/9781408855652.jpg",
		Price: 350,
		Title: "Harry Potter and the Philosopher's Stone (I)"})

	db.Create(&Book{Id: "9781408855669",
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9781/4088/9781408855652.jpg",
		Price: 350,
		Title: "Harry Potter and the Chamber of Secrets (II)"})

	db.Create(&Book{Id: "9781408855676",
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9781/4088/9781408855676.jpg",
		Price: 340,
		Title: "Harry Potter and the Prisoner of Azkaban (III)"})

	db.Create(&Book{Id: "9781408855683",
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9781/4088/9781408855683.jpg",
		Price: 360,
		Title: "Harry Potter and the Goblet of Fire (IV)"})

	db.Create(&Book{Id: "9781408855690",
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9781/4088/9781408855690.jpg",
		Price: 380,
		Title: "Harry Potter and the Order of the Phoenix (V)"})

	db.Create(&Book{Id: "9781408855706",
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9781/4088/9781408855706.jpg",
		Price: 380,
		Title: "Harry Potter and the Half-Blood Prince (VI)"})

	db.Create(&Book{
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9781/4088/9781408855713.jpg",
		Price: 400,
		Title: "Harry Potter and the Deathly Hallows (VII)",
		Id:    "9781408855713",
	})

	db.Create(&Book{
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9781/5098/9781509889969.jpg",
		Price: 270,
		Title: "The Order of the Day",
		Id:    "9781509889969",
	})

	db.Create(&Book{
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9780/2413/9780241392362.jpg",
		Price: 260,
		Title: "The Fork, the Witch, and the Worm",
		Id:    "9780241392362",
	})
	db.Create(&Book{
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9781/4711/9781471128387.jpg",
		Price: 220,
		Title: "The Dakota Winters",
		Id:    "9781471128387",
	})
	db.Create(&Book{
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9781/4732/9781473225046.jpg",
		Price: 320,
		Title: "Mistborn: Secret History",
		Id:    "9781473225046",
	})
	db.Create(&Book{
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9781/7847/9781784742539.jpg",
		Price: 280,
		Title: "Professor Chandra Follows His Bliss",
		Id:    "9781784742539",
	})
	db.Create(&Book{
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9781/9257/9781925773477.jpg",
		Price: 180,
		Title: "The Rosie Result",
		Id:    "9781925773477",
	})
	db.Create(&Book{
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9780/5713/9780571351749.jpg",
		Price: 275,
		Title: "Come Rain or Come Shine",
		Id:    "9780571351749",
	})
	db.Create(&Book{
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9781/7606/9781760632021.jpg",
		Price: 310,
		Title: "The Illumination of Ursula Flight",
		Id:    "9781760632021",
	})
	db.Create(&Book{
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9781/4736/9781473667815.jpg",
		Price: 290,
		Title: "A Shout in the Ruins",
		Id:    "9781473667815",
	})
	db.Create(&Book{
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9780/8578/9780857844774.jpg",
		Price: 260,
		Title: "The Ocean Book",
		Id:    "9780857844774",
	})
	db.Create(&Book{
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9780/1419/9780141979090.jpg",
		Price: 220,
		Title: "Enlightenment Now",
		Id:    "9780141979090",
	})
	db.Create(&Book{
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9780/2413/9780241352465.jpg",
		Price: 190,
		Title: "The New Complete Book of Self-Sufficiency",
		Id:    "9780241352465",
	})
	db.Create(&Book{
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9780/2419/9780241975367.jpg",
		Price: 230,
		Title: "The Monk of Mokha",
		Id:    "9780241975367",
	})
	db.Create(&Book{
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9781/7847/9781784706883.jpg",
		Price: 300,
		Title: "The Language of Kindness",
		Id:    "9781784706883",
	})
	db.Create(&Book{
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9780/1419/9780141987392.jpg",
		Price: 250,
		Title: "Elastic",
		Id:    "9780141987392",
	})
	db.Create(&Book{
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9781/7887/9781788701495.jpg",
		Price: 230,
		Title: "Death Row: The Final Minutes",
		Id:    "9781788701495",
	})
	db.Create(&Book{
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9781/4071/9781407192123.jpg",
		Price: 310,
		Title: "Dog Man 4: Dog Man and Cat Kid",
		Id:    "9781407192123",
	})
	db.Create(&Book{
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9781/4449/9781444941135.jpg",
		Price: 180,
		Title: "The Colour of the Sun",
		Id:    "9781444941135",
	})
	db.Create(&Book{
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9781/6014/9781601427991.jpg",
		Price: 400,
		Title: "The Minimalist Home: A Room-By-Room Guide to a Decluttered, Refocused Life",
		Id:    "9781601427991",
	})
	db.Create(&Book{
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9781/4736/9781473659711.jpg",
		Price: 120,
		Title: "How to Survive the End of the World ",
		Id:    "9781473659711",
	})
	db.Create(&Book{
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9781/5098/9781509809950.jpg",
		Price: 160,
		Title: "Solve For Happy",
		Id:    "9781509809950",
	})
	db.Create(&Book{
		Cover: "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9781/4736/9781473634176.jpg",
		Price: 345,
		Title: "The Confidence Project",
		Id:    "9781473634176",
	})

	return db
}
