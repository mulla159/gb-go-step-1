package main

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	Port        int    `envconfig:"PORT" default:"8081" required:"true"`
	DBUrl       string `envconfig:"DB_URL" default:"postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable" required:"true"`
	JaegerUrl   string `envconfig:"JAEGER_URL" required:"true"`
	SentryUrl   string `envconfig:"SENTRY_URL" required:"true"`
	KafkaBroker string `envconfig:"KAFKA_BROKER" required:"true"`
	SomeAppId   string `envconfig:"SOME_APP_ID" required:"true"`
	SomeAppKey  string `envconfig:"SOME_APP_KEY" required:"true"`
}

func main() {
	var conf AppConfig
	err := envconfig.Process("", &conf)

	if err != nil {
		fmt.Errorf("can't process the config: %w", err)
	}

	fmt.Println("Configs: ", conf)
}
