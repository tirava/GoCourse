package main

import (
	"fmt"
	"math"
)

func main() {
	c := math.Sqrt(88)
	fmt.Printf("%.1f\n", c)
	fmt.Println("Address", &c)
	//111
	fmt.Println(len("Hello World"))
	fmt.Println(" Hello World"[0])
	fmt.Println("Hello " + "World")
}
