package main

import (
	"flag"
	"log"
)

var (
	name string
)

func main() {
	flag.StringVar(&name, "name", "config.json", "profile name")
	flag.Parse()
	if name == "" {
		log.Println("error: name is nil")
		return
	}
}
