package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args[0] is the program path — real args start at [1]
	city := "the world"

	if len(os.Args) > 1 {
		city = os.Args[1]
	}

	fmt.Printf("weather for %s\n", city)
}
