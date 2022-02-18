package main

import (
	"broker/src/start"
	"fmt"
	"log"
)

func main() {
	fmt.Println("startintg application")

	errs := start.StartApplication()
	for {
		err := <-*errs
		log.Fatal("error - exiting application", err)
	}
}
