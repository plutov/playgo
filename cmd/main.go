package main

import (
	"fmt"
	"github.com/plutov/playgo"
	"os"
)

func main() {
	url, err := playgo.ShareAndOpen()
	if err != nil {
		fmt.Println("Error: " + err.Error())
		fmt.Println("USAGE: playgo [FILE]")
		os.Exit(1)
	}

	fmt.Println(url)
}
