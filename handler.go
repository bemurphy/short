package short

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	store Store
	*mux.Router
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	/* TODO Yeah...better 404 it instead */
	val, _ := h.store.Get(vars["key"])
	http.Redirect(w, r, val, http.StatusFound)
}

func NewHandler(s Store) *Handler {
	h := &Handler{
		store:  s,
		Router: mux.NewRouter(),
	}

	h.Handle("/{key}", http.HandlerFunc(h.get)).Methods("GET")

	return h
}
