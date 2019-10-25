package main

import (
	"fmt"
	"log"

	uarand "git.hydra-project.io/banks/uarand"
)

func main() {

	userAgent, err := uarand.GetRandomUserAgent()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(userAgent)
}
