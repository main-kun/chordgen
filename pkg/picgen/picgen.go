package picgen

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

const (
	whiteKeyWidth  = 48
	blackKeyWidth  = 28
	imageWidth     = whiteKeyWidth * 14
	imageHeight    = 200
	blackKeyHeight = 132
)

//var steps = [24]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}

func arrayContains(arr []int, item int) bool {
	for _, a := range arr {
		if a == item {
			return true
		}
	}
	return false

}

func DrawChord(highlights []int, filename string) error {
	fullFile, err := os.Create(filename)
	whiteKeysWithoutBlack := []int{2, 6, 9, 13}
	whiteKeyIndexStutter := []int{4, 11, 16}
	blackKeyIndexJumps := []int{3, 10, 15}
	if err != nil {
		return err
	}
	fullImage := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))
	highlightColor := color.RGBA{R: 230, G: 57, B: 70, A: 255}
	// Create a white key
	whiteKey := image.NewRGBA(image.Rect(0, 0, whiteKeyWidth, imageHeight))
	draw.Draw(whiteKey, whiteKey.Bounds(), &image.Uniform{C: image.Black}, image.Point{}, draw.Src)
	draw.Draw(whiteKey, image.Rect(1, 1, whiteKeyWidth-1, imageHeight-1), &image.Uniform{C: image.White}, image.Point{}, draw.Src)
	// Create a black key
	blackKey := image.NewRGBA(image.Rect(0, 0, blackKeyWidth, blackKeyHeight))
	draw.Draw(blackKey, blackKey.Bounds(), &image.Uniform{C: image.Black}, image.Point{}, draw.Src)
	// Create a highlighted white key
	whiteKeyHighlight := image.NewRGBA(image.Rect(0, 0, whiteKeyWidth, imageHeight))
	draw.Draw(whiteKeyHighlight, whiteKeyHighlight.Bounds(), &image.Uniform{C: image.Black}, image.Point{}, draw.Src)
	draw.Draw(whiteKeyHighlight, image.Rect(1, 1, whiteKeyWidth-1, imageHeight-1), &image.Uniform{C: highlightColor}, image.Point{}, draw.Src)
	// Create a highlighted black key
	blackKeyHighlight := image.NewRGBA(image.Rect(0, 0, blackKeyWidth, blackKeyHeight))
	draw.Draw(blackKeyHighlight, blackKey.Bounds(), &image.Uniform{C: image.Black}, image.Point{}, draw.Src)
	draw.Draw(blackKeyHighlight, image.Rect(1, 1, blackKeyWidth-1, blackKeyHeight-1), &image.Uniform{C: highlightColor}, image.Point{}, draw.Src)

	whiteKeyIndex := 0
	for i := 0; i < 14; i++ {
		whiteKeyPosition := i * whiteKeyWidth
		var keyToDraw *image.RGBA
		if arrayContains(highlights, whiteKeyIndex) {
			keyToDraw = whiteKeyHighlight
		} else {
			keyToDraw = whiteKey
		}
		draw.Draw(fullImage, image.Rect(whiteKeyPosition, 0, imageWidth, imageHeight), keyToDraw, image.Point{}, draw.Src)
		if arrayContains(whiteKeyIndexStutter, whiteKeyIndex) {
			whiteKeyIndex += 1
		} else {
			whiteKeyIndex += 2
		}
	}
	blackKeyIndex := 1
	for i := 0; i < 14; i++ {
		whiteKeyPosition := i * whiteKeyWidth
		blackKeyPosition := whiteKeyPosition + blackKeyWidth + blackKeyWidth/6
		if !arrayContains(whiteKeysWithoutBlack, i) {
			var keyToDraw *image.RGBA
			if arrayContains(highlights, blackKeyIndex) {
				keyToDraw = blackKeyHighlight
			} else {
				keyToDraw = blackKey
			}
			draw.Draw(fullImage, image.Rect(blackKeyPosition, 0, imageWidth, imageHeight), keyToDraw, image.Point{}, draw.Src)
			if arrayContains(blackKeyIndexJumps, blackKeyIndex) {
				blackKeyIndex += 3
			} else {
				blackKeyIndex += 2
			}
		}
	}
	err = png.Encode(fullFile, fullImage)
	if err != nil {
		return err
	}
	return nil
}
