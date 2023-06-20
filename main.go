package main

import (
	"fmt"

	b "github.com/un-versed/example-runtime-caller/package-b"
)

func main() {
	// Call package-b
	err := b.FromB(true)
	if err != nil {
		fmt.Println(err)
	}
}
