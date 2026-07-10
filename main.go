// Command weather is the Project P01 starter for the "Learn Go by Building" curriculum.
//
// It runs today: `go run . London` prints the city back. Your job is to grow it
// into a real weather CLI by following Project P01, step by step. The TODOs below
// map to the P01 build steps — fill them in yourself; that's the whole point.
package main

import (
	"fmt"
	"os"
)

func main() {
	// P01 · Step 02 — read the city from the command line.
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: weather <city>")
		os.Exit(1) // non-zero exit = failure, the Unix way
	}
	city := os.Args[1]

	// TODO P01 · Step 03 — fetch(ctx, city): call a weather API (e.g. wttr.in),
	//      handle the error, and defer resp.Body.Close().
	// TODO P01 · Step 04 — decode the JSON into a struct using struct tags.
	// TODO P01 · Step 05 — format the output; os.Exit(1) on any failure.
	fmt.Printf("weather for %s (TODO: fetch real data — see Project P01)\n", city)
}
