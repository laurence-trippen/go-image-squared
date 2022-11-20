package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ttacon/chalk"
)

// Usage:
// go-image-squared -path=example.png -bg=#333

func main() {
	fmt.Println(chalk.Green, "Go Image Squared v1.0", chalk.Reset)

	path := flag.String("path", "", "Path to image.")
	bg := flag.String("bg", "#000000", "Background color")

	flag.Parse()

	fmt.Println(*path)
	fmt.Println(*bg)

	if *path == "" {
		fmt.Println(chalk.Red, "Image path not specified!")
		os.Exit(0)
	}
}
