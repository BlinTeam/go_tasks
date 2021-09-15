//Напишите программу, принимающая на вход число N (N ≥ 4), а затем N целых чисел-элементов среза. На вывод нужно
//подать 4-ый (3 по индексу) элемент данного среза.

package main
import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	a := make([]int, N, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Println(a[3])
}