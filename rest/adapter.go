package rest

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Adapter struct {
	r *mux.Router
}

func NewAdapter() Adapter {
	return Adapter{mux.NewRouter()}
}

func (a Adapter) ListenAndServe() {
	log.Printf("Starting migomeal REST server on port 8080.")
	// todo add error handling
	_ = http.ListenAndServe(":8080", a.r)
}

func (a Adapter) HandleFunc(path string, f func(w http.ResponseWriter, r *http.Request)) *mux.Route {
	return a.r.NewRoute().Path(path).HandlerFunc(f)
}
