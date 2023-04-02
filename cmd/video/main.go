package main

import (
	"log"
	"os"
	"path/filepath"
	"vid"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	path := filepath.Join(cwd, "clip.mp4")

	if err = vid.PlayVideo(path, 130); err != nil {
		log.Fatal(err)
	}
}
