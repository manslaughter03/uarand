package main

import (
	"fmt"
	"log"

	uarand "github.com/manslaughter03/uarand"
)

func main() {

	userAgent, err := uarand.GetRandomUserAgent()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(userAgent)
}
