package main

import (
	"broker/src/start"
	"fmt"
)

func main() {
	fmt.Println("startintg application")

	errs := start.StartApplication()
	for {
		err := <-*errs
		fmt.Println(err)
	}
}
