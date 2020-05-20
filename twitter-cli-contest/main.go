package main

import (
	"log"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

func main() {
	anaconda.SetConsumerKey(os.Getenv("KEY"))
	anaconda.SetConsumerSecret(os.Getenv("SECRET"))
	api := anaconda.NewTwitterApi(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))

	searchResult, _ := api.GetSearch("ukvcas", nil)
	for _, tweet := range searchResult.Statuses {
		id := tweet.Id
		tweets, err := api.GetRetweets(id, nil)
		if err != nil {
			log.Panic(err)
		}
		file, err := os.OpenFile("data", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0640)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		for _, t := range tweets {
			_, err = file.WriteString(t.Text)
			if err != nil {
				log.Panic(err)
			}
		}
	}
}
