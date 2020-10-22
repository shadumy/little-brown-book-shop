package app

import (
	"go-with-compose/app/handler"
	"go-with-compose/mux"
	"net/http"
)

type App struct {
	Router *mux.Router
}

// Initialize initializes the app with predefined configuration
func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	}
}

// setRouters sets the all required routers
func (a *App) setRouters() {
	a.Get("/", a.handleRequest(handler.HomePage))
	a.Get("/getAddress", a.handleRequest(handler.GetAllBooks))
	a.Post("/discount", a.handleRequest(handler.GetDiscount))
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

type RequestHandlerFunction func(w http.ResponseWriter, r *http.Request)
