package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bemurphy/short"
	"github.com/gorilla/mux"
)

type Handler struct {
	store short.Store
	*mux.Router
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	/* TODO Yeah...better 404 it instead */
	val, _ := h.store.Get(vars["key"])
	http.Redirect(w, r, val, http.StatusFound)
}

func main() {
	url := "http://www.reddit.com"
	s, err := short.Shorten(url)
	if err != nil {
		log.Fatal(err)
	}

	h := &Handler{
		store:  short.NewMemoryStore(),
		Router: mux.NewRouter(),
	}

	h.store.Set(s, url)
	fmt.Printf("http://127.0.0.1:8080/%s\n", s)

	h.Handle("/{key}", http.HandlerFunc(h.get)).Methods("GET")
	http.ListenAndServe(":8080", h)
}
