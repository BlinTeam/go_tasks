//Напишите программу, которая в последовательности чисел находит сумму двузначных чисел, кратных 8.
//Программа в первой строке получает на вход число n - количество чисел в последовательности, во второй строке -- n чисел,
//входящих в данную последовательность.

package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	var sum = 0
	var c int
	for i := 0; i < a; i++ {
		fmt.Scan(&c)
		if c >= 10 && c <= 99 && c%8 == 0 {
			sum +=  c
		}
	}
	fmt.Println(sum)
}
