//Дана последовательность, состоящая из целых чисел. Напишите программу,
//которая подсчитывает количество положительных чисел среди элементов последовательности.
//
//Входные данные
//
//Сначала задано число N — количество элементов в последовательности (1≤N≤100). Далее через пробел записаны N чисел —
//элементы последовательности. Последовательность состоит из целых чисел.
//
//Выходные данные
//
//Необходимо вывести единственное число - количество положительных элементов в последовательности.

package main
import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	array := make([]int, N, N)
	c := 0
	for i := 0; i < N; i++ {
		fmt.Scan(&array[i])
		if array[i] > 0 {
			c++
		}
	}
	fmt.Println(c)
}