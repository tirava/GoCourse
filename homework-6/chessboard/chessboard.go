// Homework-6: Standard library - Part 2
// Exercise 2 - Chess Board
// Author: Eugene Klimov
// Date: 14 may 2019
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

const (
	rectX    = 1000 // board size
	rectY    = 1000
	fileName = "chessboard.png"
)

// Draw chessboard without figures and letters
func main() {
	brownOut := color.RGBA{R: 110, G: 50, B: 50, A: 255}
	brownDark := color.RGBA{R: 45, G: 25, B: 20, A: 255}
	brownLight := color.RGBA{R: 250, G: 200, B: 145, A: 255}

	dX, dY := rectX/10, rectY/10 // square side length

	rectImg := image.NewRGBA(image.Rect(0, 0, rectX, rectY))
	draw.Draw(rectImg, rectImg.Bounds(), &image.Uniform{C: brownOut}, image.ZP, draw.Src)
	draw.Draw(rectImg, image.Rect(dX, dY, rectX-dX, rectY-dY), &image.Uniform{C: brownDark}, image.ZP, draw.Src)

	// Draw squares
	var k int
	for j := 1; j < 9; j++ {
		for i := 1; i < 9; i += 2 {
			if j%2 == 0 { // even lines
				k = i + 1
			} else { // odd lines
				k = i
			}
			draw.Draw(rectImg, image.Rect(k*dX, j*dY, k*dX+dX, j*dY+dY),
				&image.Uniform{C: brownLight}, image.ZP, draw.Src)
		}
	}

	// Write file
	file, err := os.Create(fileName)
	check(err, "fatal", "Failed create file:")
	defer file.Close()

	err = png.Encode(file, rectImg)
	check(err, "fatal", "Failed render image file:")

	fmt.Println("Image file", fileName, "written successfully!")
}

// check is errors helper
func check(err error, sType, sMessage string) {
	if err != nil {
		switch sType {
		case "fatal":
			log.Fatalln(sMessage, err)
		default:
			log.Println(sMessage, err)
		}
	}
}
