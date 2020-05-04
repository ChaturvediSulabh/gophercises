package main

import (
	"log"
	"net/http"
)

var url string = "https://www.digitalocean.com/community/tutorials/importing-packages-in-go"

func main() {
	http.HandleFunc("/importing-packages-in-go", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/importing-packages-in-go" {
		http.Redirect(w, r, url, http.StatusFound)
	}
}
