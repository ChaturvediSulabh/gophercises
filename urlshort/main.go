package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/streadway/amqp"
)

//Config is slice of struct of route and their respective Redirects
type Config struct {
	URLMap map[string]string
}

func main() {
	var c Config
	conf, err := c.amqpHandler()
	failOnError(err)
	handler := redirectHandler(conf, fallback())
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func redirectHandler(config *Config, fallback http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		urlExists := false
		for k, v := range config.URLMap {
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

func (conf *Config) amqpHandler() (*Config, error) {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_CONNECTION_STRING"))
	failOnError(err)
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err)
	defer ch.Close()

	q, err := ch.QueueDeclare(os.Getenv("RABBITMQ_URLSHORT_QUEUE_NAME"), false, false, false, false, nil)
	failOnError(err)

	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	failOnError(err)

	go func() {
		for msg := range msgs {
			fmt.Println("message received: " + string(msg.Body))
			slOfMsg := strings.Split(string(msg.Body), ";")
			conf.URLMap = make(map[string]string)
			conf.URLMap[slOfMsg[0]] = slOfMsg[1]
		}
	}()
	fmt.Println(conf)
	return conf, nil
}

func failOnError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
