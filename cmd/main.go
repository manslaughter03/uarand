package main

import (
	"fmt"
	"log"

	uarand "git.hydra-project.io/banks/user-agent-random"
)

func main() {

	userAgent, err := uarand.GetRandomUserAgent()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(userAgent)
}
