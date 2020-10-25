# little-brown-book-shop

API sample data

https://json-bin.netlify.app/books.json

to remote Access

http://209.97.167.166:9567/

# Installation and Run in local machine

- Clone Project to any directory (no need to put in go directory in your machine because we run on docker)

```bash
git clone https://github.com/shadumy/little-brown-book-shop.git
```

- change directory to project directory and then create .env file from template (.env.template)

```bash
cd little-brown-book-shop
cp .env.template .env
```

- edit .env file with your data. project will use these as parameters. You can see an each field input guide in parenthesis.

```bash
nano .env
```

for _DB_HOST_, please find your local ip address by this command. Please do not use localhost or 127.0.0.1 because it might cause a connection problem between api docker and db docker.

```bash
ipconfig getifaddr en0
```

This is an example of completed .env
_note_ do not just copy-paste and run it, it has to be your desired security data.

```
DB_HOST=192.168.1.107
DB_PASSWORD=12345
DB_NAME=mainDb
DB_HOST_PORT=9568
AUTH_USER=user1
AUTH_PASSWORD=password
AUTH_KEY=my_secret_key
```

- Run Server
  _note_ The project requires docker and docker-compose installed in a machine before running.

```bash
docker-compose up -d
```

## API

#### /

- `GET` : Get a welcome greeting text

_No need any authentication for this route._

#### /token

- `POST` : get an authentication token

_Need login authentication for this route._

To get any data from this API, user has to specify token created from system. To get token, please request by send username and password in JSON format to this route. Username and password is one you specified in .env file as AUTH_USER and AUTH_PASSWORD.

_A token has expiration time, able to use within 5 minutes after generated_

Example request using curl, you can use other application like Postman or Insomnia.

```bash
# Testing with zsh in MacOS, other shell might need a bit change such as single quote or backslash
curl -H "Content-Type: application/json" -X POST -d '{"username":"{AUTH_USER}","password":"{AUTH_PASSWORD}"}' http://localhost:9567/token
```

#### /books

- `GET` : Get all list of books in system
- `POST` : Create a new book

_Need header token authentication for this route._

To get list of all books or add a new book, user has to add token in header.

Example request using curl, you can use other application like Postman or Insomnia.

```bash
# Testing with zsh in MacOS, other shell might need a bit change such as single quote or backslash

# Get all books
curl -H "Content-Type: application/json" -H "Authorization: Bearer {TOKEN}" -X GET http://localhost:9567/books

# Create a new book to system
curl -H "Content-Type: application/json" -H "Authorization: Bearer {TOKEN}" -d '{"id":"123","cover":"https://images-na.ssl-images-amazon.com/images/I/8134AkhQJgL.jpg", "price": 555, "title": "The lord of the rings"}' -X POST http://localhost:9567/books
```

Creating a new book has 1 requred field (id [string]) and 3 optional fields (cover [string], price [int], title [string])

#### /books/:id

- `GET` : Get a book data by id
- `PUT` : Update a book data by id
- `DELETE` : Delete a book by id

_Need header token authentication for this route._

To get data, update or delte a book, user has to add token in header.

Example request using curl, you can use other application like Postman or Insomnia.

```bash
# Testing with zsh in MacOS, other shell might need a bit change such as single quote or backslash

# Get a book's data
curl -H "Content-Type: application/json" -H "Authorization: Bearer {TOKEN}" -X GET http://localhost:9567/books/123

# Update a book's data
curl -H "Content-Type: application/json" -H "Authorization: Bearer {TOKEN}" -d '{"id":"123","cover":"https://images-na.ssl-images-amazon.com/images/I/8134AkhQJgL.jpg", "price": 666, "title": "The Lord of the Rings"}' -X PUT http://localhost:9567/books/123

# Delete a book from system
curl -H "Content-Type: application/json" -H "Authorization: Bearer {TOKEN}" -X DELETE http://localhost:9567/books/123
```

#### /discount

- `POST` : Get a net discount from books data

_Need header token authentication for this route._

To get net discount from list of books data, user has to add token in header.

Example request using curl, you can use other application like Postman or Insomnia.

```bash
# Testing with zsh in MacOS, other shell might need a bit change such as single quote or backslash

# Get net discount
curl -H "Content-Type: application/json" -H "Authorization: Bearer {TOKEN}" -d '{"books": [{"cover": "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9781/4088/9781408855652.jpg","price": 350,"title": "Harry Potter and the Philosopher Stone (I)","id": "9781408855652"}, {"cover": "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9781/4088/9781408855652.jpg", "price": 350, "title": "Harry Potter and the Philosopher Stone (I)","id": "9781408855652"}, {"cover": "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9781/4088/9781408855652.jpg", "price": 350,
"title": "Harry Potter and the Chamber of Secrets (II)",
"id": "9781408855669"},{"cover": "https://d1w7fb2mkkr3kw.cloudfront.net/assets/images/book/mid/9781/4088/9781408855652.jpg","price": 260,"title": "The Fork, the Witch, and the Worm",
"id": "9780241392362"}]}' -X POST http://localhost:9567/discount
```

JSON detail in body has to be in format :

```
{"books": [{ {book1}, {book2}, {book3} ...}]}
```

Each book (book1, book2 ,....) has to be in format :

```
{"id": "", "cover": "", "title": "", price: 0}]}
```

For current version, system does not check other fields except id and price to calculate a discount. User can let cover and title as an empty string
