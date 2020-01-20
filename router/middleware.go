package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/riphidon/evo/config"
	"github.com/riphidon/evo/utils"
)

var id string

type AppHandler func(w http.ResponseWriter, r *http.Request) error
type AuthHandler func(w http.ResponseWriter, r *http.Request) error

func (fn AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("logged %v requested %v", r.RemoteAddr, r.URL)
	err := fn(w, r)
	if err != nil {
		http.Error(w, fmt.Sprintf("error : %+v\n", err), http.StatusInternalServerError)
	}
}

func (fn AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

// func favicoHandler(w http.ResponseWriter, r *http.Request) error {
// 	http.Redirect(w, r, "/static/assets/images/logo.ico", http.StatusSeeOther)
// 	return nil
// }

//midleware for routes management
func SetupRoutes(mux *mux.Router) {
	mux.Handle("/static/", config.FileSystem)
	mux.Handle("/", AppHandler(getItems)).Methods("GET")
	mux.Handle("/item/{id}", AppHandler(getItem)).Methods("GET")
	mux.Handle("/create", AppHandler(createItem)).Methods("POST")
	mux.Handle("/update/{id}", AppHandler(updateItem)).Methods("PUT")
	mux.Handle("/delete/{id}", AppHandler(deleteItem)).Methods("DELETE")
	// mux.Handle("/favicon.ico", AppHandler(favicoHandler))

}

// mux.Handle("/", routes.Home)
// mux.Handle("/login", routes.Login)
// mux.Handle("/register", routes.Register)
// mux.Handle("/profile/", routes.Profile)
