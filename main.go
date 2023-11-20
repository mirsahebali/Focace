package main

import (
	"fmt"

	"mirsahebali/focace/api"
)

func main() {
	err := api.Run()
	if err != nil {
		fmt.Println(err)
		fmt.Println("cannot start server")
	}
}
