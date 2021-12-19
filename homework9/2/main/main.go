package main

import (
	"flag"
	"gb-go-step-1/homework9/2/config"
)

func main() {
	conf := config.AppConfig{}

	var fileName = flag.String("fn", "", "File name (path)")
	flag.Parse()

	conf.Get(*fileName)

	conf.Print()
}
