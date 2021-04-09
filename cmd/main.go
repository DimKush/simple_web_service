package main

import (
	"log"
	process "simple_web_service/internal"
)

func main() {
	err := process.ResponceSize("http://google.com")
	if err != nil {
		log.Fatal(err)
	}
}
