//Напишите программу, которая выводит квадраты натуральных чисел от 1 до 10. Квадрат каждого числа должен выводится в новой строке.

package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println((i + 1) * (i + 1))
	}
}