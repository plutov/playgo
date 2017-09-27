package main

import (
	"fmt"
	"os"

	"github.com/plutov/playgo"
)

func main() {
	url, err := playgo.ShareAndOpen()
	if err != nil {
		fmt.Printf("Error: %s\nUSAGE: playgo [FILE]\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("%s (copied to clipboard)\n", url)
}
