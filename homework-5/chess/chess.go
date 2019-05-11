// Homework-4: Methods and Interfaces
// Author: Eugene Klimov
// Date: 11 may 2019
package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Point struct {
	x, y int
}

// delta needs for calc horse XY coordinates
var delta = [8][2]int{{1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}, {-2, -1}, {-2, 1}, {-1, 2}}

// Exercise #4 - horse steps
func main() {

	var step string
	var sX, sY string

	for {
		fmt.Printf("\nEnter position of the horse [for example e2] or 'exit': ")
		fmt.Scanln(&step)
		step = strings.ToLower(step)
		if step == "exit" {
			os.Exit(0)
		}
		sX, sY = step[:1], step[1:]

		if sX < "a" || sX > "h" || sY < "1" || sY > "8" {
			fmt.Println("Error, repeat please...")
			continue
		}

		// x & y - convert horse coordinates from console chess to integer
		x, y := getXYFromChess(sX, sY)

		// create horse
		horse := Point{}
		horse.setHorseXY(x, y)

		//get possible steps & check calc performance
		start := time.Now()
		points := horse.getHorseSteps()
		end := time.Now()
		log.Println("Calculation of moves was performed for", end.Sub(start))

		// print steps
		fmt.Printf("Possible moves are: ")
		for _, p := range points {
			sX, sY = getChessFromXY(p.x, p.y)
			fmt.Printf("%s ", sX+sY)
		}
		fmt.Println("")
	}
}

// getHorseSteps gets possible steps
func (p Point) getHorseSteps() (steps []Point) {
	var step Point
	var x, y int

	steps = make([]Point, 0, 8)

	// check by 'simple circle' method that more easier and faster than 'mirror circle'
	for i := 0; i < 8; i++ {
		x, y = p.x+delta[i][0], p.y+delta[i][1]
		if 0 > x || y < 0 || 7 < x || y > 7 { // hmm
			continue // if outbound from the desk
		}
		step = Point{x, y}
		steps = append(steps, step)
	}

	return steps
}

// getChessFromXY
func getChessFromXY(x, y int) (sX, sY string) {
	sX, sY = string(byte(x+97)), string(byte(y+49)) // convert int to char via ascii
	return sX, sY
}

// getXYFromChess
func getXYFromChess(sX, sY string) (x, y int) {
	bX, bY := []byte(sX), []byte(sY)
	x, y = int(bX[0])-97, int(bY[0])-49 // convert char to int via ascii
	return x, y
}

// setHorseXY sets x & y
func (p *Point) setHorseXY(x, y int) {
	p.x, p.y = x, y
}
