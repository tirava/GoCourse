package main

import (
	"fmt"
	"log"
	"os"
)

// check упрощает код при множественных однотипных обработках ошибок
func check(e error) {
	// Логгируем в консоль и выходим
	if e != nil {
		log.Fatal(e)
	}
}

func main() {

	file, err := os.Open("fileread.go")
	//if err != nil {
	//    return
	//}
	check(err)
	defer file.Close()

	// getting size of file
	stat, err := file.Stat()
	// if err != nil {
	//    return
	//}
	check(err)

	// reading file
	bs := make([]byte, stat.Size())
	// Здесь тоже можно сделать сильно короче, т.к. получаем только один аргумент (err), но хуже читаемость
	//_, err = file.Read(bs)
	//if err != nil {
	//    return
	//}
	_, err = file.Read(bs)
	//if err != nil {
	//    return
	//}
	check(err) // err из Read еще может возвращать EOF, но здесь лучше не усложнять

	fmt.Println(string(bs))
}
