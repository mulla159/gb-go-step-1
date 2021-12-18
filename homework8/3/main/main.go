package main

import (
	"gb-go-step-1/homework8/3/config"
)

func main() {
	var conf config.AppConfig

	conf.Get()

	conf.Print()
}
