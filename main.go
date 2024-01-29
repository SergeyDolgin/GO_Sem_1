package main

import (
	"fmt"
	"strings"
)

func main() {
	var myfile string

	fmt.Print("Введите имя файла в формате имя.расширение (***.***)")
	fmt.Scan(&myfile)
	// вывод на экран имени файла
	myfilereverse := Reverse(myfile)
	fmt.Println("Обратное название", myfilereverse)
	name := strings.LastIndex(myfilereverse, ".")
	if name == -1 {
		fmt.Println("Неверно задан путь к файлу. Нет имени или расшиирения.")
		return
	}
	nme := myfilereverse[name:]
	fmt.Println("filename:", myfilereverse, "равно", Reverse(nme))
	// вывод на экран расширения файла
	extension := strings.LastIndex(myfile, ".")
	ext := myfile[extension:]
	fmt.Println("extension:", myfile, "равно", ext)
}

// функция разворота строки
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
