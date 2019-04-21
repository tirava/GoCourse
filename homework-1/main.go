package main

import (
	"fmt"
	"log"
	"strconv"
)

const USD2RUB = 64

func main() {
	var sRub string
	fmt.Println("\nУкажите количество рублей: ")
	_, _ = fmt.Scanln(&sRub)
	fRub, err := strconv.ParseFloat(sRub, 2)
	if err != nil {
		log.Fatal(err)
	}
	fUsd := convertRub2Usd(fRub)
	fmt.Printf("\nКонвертация в доллары:\n%.2f\n", fUsd)
}

func convertRub2Usd(fRub float64) float64 {
	fUsd := fRub / USD2RUB
	return fUsd
}
