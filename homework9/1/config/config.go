package config

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
)

type CustomURL struct {
	Value *url.URL
}

func (cu *CustomURL) UnmarshalJSON(data []byte) error {
	var urlString string

	err := json.Unmarshal(data, &urlString)
	if err != nil {
		return err
	}

	u, err := url.ParseRequestURI(urlString)
	cu.Value = u
	return err
}

type AppConfig struct {
	Port        int       `json:"port"`
	DbURL       CustomURL `json:"db_url"`
	JaegerURL   CustomURL `json:"jaeger_url"`
	SentryURL   CustomURL `json:"sentry_url"`
	KafkaBroker string    `json:"kafka_broker"`
	SomeAppID   string    `json:"some_app_id"`
	SomeAppKey  string    `json:"some_app_key"`
}

type AppConfig1 struct {
	Port        int
	DbURL       string
	JaegerURL   string
	SentryURL   string
	KafkaBroker string
	SomeAppID   string
	SomeAppKey  string
}

func (conf *AppConfig) Get(fileName string) {
	if fileName == "" {
		log.Fatalf("Не указано имя файла конфигурации")
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Не могу открыть файл: %v", err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatalf("Не могу закрыть файл: %v\n", err)
		}
	}()

	err = json.NewDecoder(file).Decode(&conf)
	if err != nil {
		log.Fatalf("Не могу декодировать json-файл в структуру: %v\n", err)
	}
}

func (conf *AppConfig) Print() {
	fmt.Println("Configs:")
	fmt.Println("Port: ", conf.Port)
	fmt.Println("DBUrl: ", conf.DbURL)
	fmt.Println("JaegerUrl: ", conf.JaegerURL)
	fmt.Println("SentryUrl: ", conf.SentryURL)
	fmt.Println("KafkaBroker: ", conf.KafkaBroker)
	fmt.Println("SomeAppId: ", conf.SomeAppID)
	fmt.Println("SomeAppKey: ", conf.SomeAppKey)
}
