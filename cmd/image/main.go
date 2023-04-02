package main

import (
	"github.com/kryptamine/vid"
	"log"
	"os"
	"path/filepath"
)

const width = 60

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	path := filepath.Join(cwd, "apple.png")

	if err = vid.PrintImage(path, width); err != nil {
		log.Fatal(err)
	}
}
