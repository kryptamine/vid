package main

import (
	"github.com/kryptamine/vid"
	"log"
	"os"
	"path/filepath"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	path := filepath.Join(cwd, "clip.mp4")

	if err = vid.PlayVideo(path, 100); err != nil {
		log.Fatal(err)
	}
}
