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
		urlExists := false
		for k, v := range urlMap {
			if r.URL.Path == k {
				urlExists = true
				http.Redirect(w, r, v, http.StatusFound)
			}
		}
		if urlExists == false {
			fallback.ServeHTTP(w, r)
		}
	})
}

func fallback() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	})
}
