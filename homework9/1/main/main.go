package main

import (
	"flag"
	"gb-go-step-1/homework9/1/config"
)

func main() {
	var conf config.AppConfig

	var fileName = flag.String("fn", "", "File name (path)")
	flag.Parse()

	conf.Get(*fileName)

	conf.Print()
}
