package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

func main() {
	config := flag.String("urlConfig", "./config.yaml", "Config File containing all routes and their associated urls. As, for route urlshort-godoc, its redirect is https://godoc.org/github.com/gophercises/urlshort")
	flag.Parse()

	yamlHandler, err := yamlHandler(config, fallback())
	if err != nil {
		log.Panic(err)
	}
	log.Fatal(http.ListenAndServe(":8080", yamlHandler))
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

func yamlHandler(config *string, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYAML, err := parseYAML(*config)
	if err != nil {
		log.Panicln(err)
	}
	urlMap := buildMap(parsedYAML)
	return mapHandler(urlMap, fallback), nil
}

type yamlConfig []struct {
	Route string `yaml:"route"`
	URL   string `yaml:"url"`
}

func parseYAML(filePath string) (yamlConfig, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var conf yamlConfig
	err = yaml.Unmarshal(content, &conf)
	return conf, nil
}

func buildMap(data yamlConfig) map[string]string {
	urlMap := make(map[string]string)
	for _, v := range data {
		urlMap[v.Route] = v.URL
	}
	return urlMap
}
