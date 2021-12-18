package config

import (
	"fmt"
	"log"
	"net/url"

	"github.com/kelseyhightower/envconfig"
)

// https://github.com/kelseyhightower/envconfig/blob/master/envconfig_test.go
type CustomURL struct {
	Value *url.URL
}

func (cu *CustomURL) UnmarshalBinary(data []byte) error {
	u, err := url.ParseRequestURI(string(data))
	cu.Value = u
	return err
}

type AppConfig struct {
	Port        int       `envconfig:"PORT" default:"8081" required:"true"`
	DBUrl       string    `envconfig:"DB_URL" default:"postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable" required:"true"`
	JaegerUrl   CustomURL `envconfig:"JAEGER_URL" required:"true"`
	SentryUrl   CustomURL `envconfig:"SENTRY_URL" required:"true"`
	KafkaBroker string    `envconfig:"KAFKA_BROKER" required:"true"`
	SomeAppId   string    `envconfig:"SOME_APP_ID" required:"true"`
	SomeAppKey  string    `envconfig:"SOME_APP_KEY" required:"true"`
}

func (conf *AppConfig) Get() {
	err := envconfig.Process("SENTRY_URL", conf)

	if err != nil {
		log.Fatal("can't process the config: %w", err)
	}

	fmt.Println(conf.SentryUrl)
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
