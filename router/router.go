package router

import (
	"github.com/gorilla/mux"
)

// InitRouter instantiate a mux type router
func InitRouter() *mux.Router {
	r := mux.NewRouter()
	mux.CORSMethodMiddleware(r)
	return r
}
