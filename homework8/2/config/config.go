package config

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
)

type AppConfig struct {
	Port      int
	JaegerUrl url.URL
	SentryUrl url.URL
}

func (conf *AppConfig) Get() {
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port is not set")
	}

	port, err := strconv.ParseInt(portString, 10, 64)

	if err != nil {
		log.Fatal("Port is wrong: ", err)
	}

	conf.Port = int(port)

	tmp, err := url.ParseRequestURI(os.Getenv("JAEGER_URL"))
	if err != nil {
		log.Fatal("JAEGER_URL is wrong", err)
	}
	conf.JaegerUrl = *tmp

	tmp, err = url.ParseRequestURI(os.Getenv("SENTRY_URL"))
	if err != nil {
		log.Fatal("SENTRY_URL is wrong", err)
	}
	conf.SentryUrl = *tmp
}

func (conf *AppConfig) Print() {
	fmt.Println("Configs:")
	fmt.Println("Port: ", conf.Port)
	fmt.Println("JaegerUrl: ", conf.JaegerUrl.String())
	fmt.Println("SentryUrl: ", conf.SentryUrl.String())
}
