package main

import (
	"log"
	"net/http"
)

func main() {
	urlMap := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := mapHandler(urlMap)
	for k := range urlMap {
		http.HandleFunc(k, mapHandler)
	}
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func mapHandler(urlMap map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for k, v := range urlMap {
			if r.URL.Path == k {
				http.Redirect(w, r, v, http.StatusFound)
			}
		}
	}
}
