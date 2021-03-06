//Последовательность состоит из натуральных чисел и завершается числом 0. Определите количество элементов этой последовательности,
//которые равны ее наибольшему элементу.
//
//Формат входных данных
//
//Вводится непустая последовательность натуральных чисел, оканчивающаяся числом 0 (само число 0 в последовательность не входит,
//а служит как признак ее окончания).
//
//Формат выходных данных
//
//Выведите ответ на задачу.

package main

import "fmt"

func main() {
	var max int  = 1
	var a int  = 1
	var sum int

	for a != 0{
		fmt.Scan(&a)
		if max < a {
			max = a
			sum = 0
		}
		if a == max {
			sum++
		}
	}
	fmt.Println(sum)
}