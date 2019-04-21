package main

import (
	"fmt"
	"math"
	"strconv"
)

const USD2RUB = 64

func main() {
	convertRub2Usd()
	triangleUtils()
}

func triangleUtils() {
	var leg1, leg2 float64
	var square, perimeter float64
	fmt.Println("\nУкажите катеты прямоугольно треугольника через пробел:")
	_, _ = fmt.Scanln(&leg1, &leg2)
	square = leg1 * leg2 / 2
	fmt.Printf("\nПлощадь треугольника:\n%.2f\n", square)
	hypo := math.Sqrt(math.Pow(leg1, 2) + math.Pow(leg2, 2))
	fmt.Printf("\nГипотенуза треугольника:\n%.2f\n", hypo)
	perimeter = leg1 + leg2 + hypo
	fmt.Printf("\nПериметр треугольника:\n%.2f\n", perimeter)
}

func convertRub2Usd() {
	var sRub string
	fmt.Println("\nУкажите количество рублей: ")
	_, _ = fmt.Scanln(&sRub)
	fRub, _ := strconv.ParseFloat(sRub, 2)
	fUsd := fRub / USD2RUB
	fmt.Printf("\nКонвертация в доллары:\n%.2f\n", fUsd)
}
