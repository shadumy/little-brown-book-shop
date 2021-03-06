package app

import (
	"fmt"
	"go-with-compose/app/handler"
	"go-with-compose/app/model"
	"go-with-compose/config"
	"go-with-compose/gorm"
	"go-with-compose/mux"
	"log"

	"net/http"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Initialize initializes the app with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
	)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.DB, w, r)
	}
}

// setRouters sets the all required routers
func (a *App) setRouters() {
	a.Get("/", a.handleRequest(handler.HomePage))
	a.Get("/books", a.handleRequest(handler.GetAllBooks))
	a.Post("/books", a.handleRequest(handler.CreateBook))
	a.Get("/books/{id}", a.handleRequest(handler.GetBook))
	a.Put("/books/{id}", a.handleRequest(handler.UpdateBook))
	a.Delete("/books/{id}", a.handleRequest(handler.DeleteBook))
	a.Post("/discount", a.handleRequest(handler.GetDiscount))
	a.Post("/token", a.handleRequest(handler.RequestToken))
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Run the app on it's router
func (a *App) Run(host string) {
	// log.Fatal(http.
	http.ListenAndServe(host, a.Router)
}

type RequestHandlerFunction func(db *gorm.DB, w http.ResponseWriter, r *http.Request)
