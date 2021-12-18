package main

import (
	"gb-go-step-1/homework8/2/config"
)

func main() {
	var conf config.AppConfig

	conf.Get()

	conf.Print()
}
