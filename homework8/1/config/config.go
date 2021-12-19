package config

import (
	"flag"
	"fmt"
	"log"
	"net/url"
)

type AppConfig struct {
	Port        int
	DBUrl       string
	JaegerUrl   url.URL
	SentryUrl   url.URL
	KafkaBroker string
	SomeAppId   string
	SomeAppKey  string
}

func (conf *AppConfig) Get() {
	var (
		port        = flag.Int("port", 8080, "port")
		DBUrl       = flag.String("db_url", "", "db_url")
		jaegerUrl   = flag.String("jaeger_url", "", "jaeger_url")
		sentryUrl   = flag.String("sentry_url", "", "sentry_url")
		kafkaBroker = flag.String("kafka_broker", "", "kafka_broker")
		someAppId   = flag.String("some_app_id", "", "some_app_id")
		someAppKey  = flag.String("some_app_key", "", "some_app_key")
	)

	flag.Parse()

	tmp, err := url.ParseRequestURI(*jaegerUrl)
	if err != nil {
		log.Fatal("jaegerUrl is wrong: ", err)
	}

	conf.JaegerUrl = *tmp

	tmp, err = url.ParseRequestURI(*sentryUrl)
	if err != nil {
		log.Fatal("sentryUrl is wrong: ", err)
	}
	conf.SentryUrl = *tmp

	conf.Port = *port
	conf.Port = *port
	conf.DBUrl = *DBUrl
	conf.KafkaBroker = *kafkaBroker
	conf.SomeAppId = *someAppId
	conf.SomeAppKey = *someAppKey
}

func (conf *AppConfig) Print() {
	fmt.Println("Configs:")
	fmt.Println("Port: ", conf.Port)
	fmt.Println("DBUrl: ", conf.DBUrl)
	fmt.Println("JaegerUrl: ", conf.JaegerUrl)
	fmt.Println("SentryUrl: ", conf.SentryUrl)
	fmt.Println("KafkaBroker: ", conf.KafkaBroker)
	fmt.Println("SomeAppId: ", conf.SomeAppId)
	fmt.Println("SomeAppKey: ", conf.SomeAppKey)
}
