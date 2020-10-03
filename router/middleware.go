package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/riphidon/evo/utils"
)

var id string

type appHandler func(w http.ResponseWriter, r *http.Request) error
type authHandler func(w http.ResponseWriter, r *http.Request) error

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("logged %v requested %v", r.RemoteAddr, r.URL)
	err := fn(w, r)
	if err != nil {
		http.Error(w, fmt.Sprintf("error : %+v\n", err), http.StatusInternalServerError)
	}
}

func (fn authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("logged %v requested %v", r.RemoteAddr, r.URL)
	cookie, errCookie := r.Cookie("session")
	if errCookie != nil {
		fmt.Printf("ErrCookie: %v\n", errCookie)
		return
	}
	fmt.Printf("cookie value is : %v\n", cookie.Value)
	value, err := utils.ReadCookieHandler(w, r)
	if err != nil {
		fmt.Printf("error reading cookie value: %v\n", err)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	id = value
	if errCookie != nil || utils.CheckCookie(cookie, cookie.Value) == false {
		fmt.Printf("error checkCookie: %v\n, %v\n", errCookie, utils.CheckCookie(cookie, cookie.Value))
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	err = fn(w, r)
	if err != nil {
		http.Error(w, fmt.Sprintf("error : %+v\n", err), http.StatusInternalServerError)
	}
}

// SetupRoutes acts as a midleware for routes management
func SetupRoutes(mux *mux.Router) {
	mux.Handle("/api/items", appHandler(getItems)).Methods("GET")
	mux.Handle("/api/item/{id}", appHandler(getItem)).Methods("GET")
	mux.Handle("/api/items", appHandler(createItem)).Methods("POST")
	mux.Handle("/api/items/{id}", appHandler(updateItem)).Methods("PUT")
	mux.Handle("/api/items/{id}", appHandler(deleteItem)).Methods("DELETE")
}
