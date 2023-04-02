package vid

import (
	"fmt"
	"golang.org/x/image/draw"
	"image"
	"os"
	"strings"
)

func Resize(img image.Image, width int) image.Image {
	if img.Bounds().Max.X <= width {
		return img
	}

	height := int(float64(width) * float64(img.Bounds().Max.Y) / float64(img.Bounds().Max.X))

	resizedImage := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.BiLinear.Scale(resizedImage, resizedImage.Bounds(), img, img.Bounds(), draw.Src, nil)

	return resizedImage
}

func load(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to load an image from: %s, %w", path, err)
	}
	defer file.Close()

	// Decode the image.
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode an image from: %w", err)
	}

	return img, nil
}

func rgbToAnsi(r, g, b uint8) int {
	if r == g && g == b {
		if r < 8 {
			return 16
		}

		if r > 248 {
			return 231
		}

		return ((int(r)-8)*24)/247 + 232
	}

	return 16 + (36*colorComponentToAnsi(r) + (6*colorComponentToAnsi(g) + colorComponentToAnsi(b)))
}

func colorComponentToAnsi(val uint8) int {
	return int(val) * 5 / 255
}

func isTrueColorSupported() bool {
	colorTerm := strings.ToLower(os.Getenv("COLORTERM"))
	return strings.Contains(colorTerm, "truecolor") || strings.Contains(colorTerm, "24bit")
}

func PrintImage(path string, width int) error {
	img, err := load(path)
	if err != nil {
		return err
	}

	img = Resize(img, width)

	printImage(img)
	return nil
}

func printImage(img image.Image) {
	isSupported := isTrueColorSupported()

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()

			r8 := uint8(r >> 8)
			g8 := uint8(g >> 8)
			b8 := uint8(b >> 8)

			if isSupported {
				fmt.Printf("\u001b[48;2;%d;%d;%dm  ", r8, g8, b8)
			} else {
				ansiCode := rgbToAnsi(
					r8,
					g8,
					b8,
				)

				fmt.Printf("\u001b[48;5;%dm  ", ansiCode)
			}
		}

		fmt.Println("\u001b[0m")
	}
}
