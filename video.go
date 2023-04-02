package vid

import (
	"fmt"
	"image"
	"time"

	v "github.com/AlexEidt/Vidio"
)

func PlayVideo(path string, width int) error {
	video, err := v.NewVideo(path)
	if err != nil {
		return err
	}

	frameCount := 0

	w := video.Width()
	h := video.Height()

	img := image.NewRGBA(image.Rect(0, 0, w, h))
	if err = video.SetFrameBuffer(img.Pix); err != nil {
		return err
	}

	// Calculate the duration to sleep between frames based on the frameCount rate
	sleepDuration := time.Duration(int64(time.Second) / int64(video.FPS()))

	// Hide the cursor
	fmt.Print("\033[?25l")

	for video.Read() {
		frameCount++

		resizedImg := Resize(img, width)

		printImage(resizedImg)

		// Move the cursor to the top-left corner of the terminal
		fmt.Print("\033[1;1H")

		// Sleep for a short duration to simulate the video frameCount rate
		time.Sleep(sleepDuration)
	}

	// Show the cursor
	fmt.Print("\033[?25h")

	return nil
}
