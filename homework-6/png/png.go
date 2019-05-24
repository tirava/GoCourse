// Homework-6: Standard library - Part 2
// Exercise 2 - PNG lines
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
	rectXY   = 300 // box size
	stepXY   = 30  // step for both axis
	fileName = "rectangle.png"
)

// Draw rectangle and lines in it (horizontal + vertical)
func main() {
	cyan := color.RGBA{R: 0, G: 255, B: 255, A: 255}
	red := color.RGBA{R: 255, G: 0, B: 0, A: 255}

	rectImg := image.NewRGBA(image.Rect(0, 0, rectXY, rectXY))
	draw.Draw(rectImg, rectImg.Bounds(), &image.Uniform{C: cyan}, image.ZP, draw.Src)

	// Draw lines
	for y := stepXY; y < rectXY; y += stepXY {
		for x := 0; x < rectXY; x++ {
			rectImg.Set(x, y, red) // horizontal
			rectImg.Set(y, x, red) // vertical
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
