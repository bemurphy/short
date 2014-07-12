package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bemurphy/short"
)

func main() {
	url := "http://www.reddit.com"
	s, err := short.Shorten(url)
	if err != nil {
		log.Fatal(err)
	}

	store := short.NewMemoryStore()

	store.Set(s, url)
	fmt.Printf("http://127.0.0.1:8080/%s\n", s)
	http.ListenAndServe(":8080", short.NewHandler(store))
}
