package main

import (
	"fmt"
	"math"
	"strconv"
)

const USD2RUB = 64
const NUM_YEAR = 5

func main() {
	convertRub2Usd() // конвертация рублей в доллары
	triangleUtils()  // рассчет площади, гипотенузы и пермитра по катетам
	calcProfit()     //рассчет суммы вклада с ежегодной капитализацией и без за NUM_YEAR лет
}

func calcProfit() {
	var sumBegin, percent string
	var fSumEnd, fSumEndCap float64

	fmt.Println("\nУкажите сумму вклада:")
	_, _ = fmt.Scanln(&sumBegin)
	fmt.Println("\nУкажите годовой процент:")
	_, _ = fmt.Scanln(&percent)
	fSumBegin, _ := strconv.ParseFloat(sumBegin, 2)
	fPercent, _ := strconv.ParseFloat(percent, 2)

	fSumEnd = fSumBegin + (fSumBegin*fPercent*NUM_YEAR)/100
	fSumEndCap = fSumBegin * math.Pow(1+(fPercent/100), NUM_YEAR)

	fmt.Printf("\nСумма вклада через %d лет будет (без капитализации):\n%.2f\n", NUM_YEAR, fSumEnd)
	fmt.Printf("\nСумма вклада через %d лет будет (с ежегодной капитализацией):\n%.2f\n", NUM_YEAR, fSumEndCap)
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

	fmt.Println("\nУкажите количество рублей для конвертации в доллары:")
	_, _ = fmt.Scanln(&sRub)
	fRub, _ := strconv.ParseFloat(sRub, 2)
	fUsd := fRub / USD2RUB
	fmt.Printf("\nКонвертация в доллары:\n%.2f\n", fUsd)
}
