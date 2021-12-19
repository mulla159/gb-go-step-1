package main

import "gb-go-step-1/homework9/1/config"

func main() {
	var conf config.AppConfig

	conf.Get("app.json")

	conf.Print()
}
