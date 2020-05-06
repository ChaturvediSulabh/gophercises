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
	mapHandler := mapHandler(urlMap, fallback())
	for k := range urlMap {
		http.HandleFunc(k, mapHandler)
	}
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func mapHandler(urlMap map[string]string, fallback http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for k, v := range urlMap {
			if r.URL.Path == k {
				http.Redirect(w, r, v, http.StatusFound)
			}
		}
		fallback.ServeHTTP(w, r)
	})
}

func fallback() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	})
}
