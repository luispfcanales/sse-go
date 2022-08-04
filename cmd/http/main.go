package main

import (
	"log"

	"github.com/luispfcanales/sse-go/cmd/http/boostrap"
)

func main() {
	err := boostrap.Run()
	if err != nil {
		log.Fatal(err)
	}
}
