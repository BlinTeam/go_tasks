//Из натурального числа удалить заданную цифру.
//
//Входные данные
//
//Вводятся натуральное число и цифра, которую нужно удалить.
//
//Выходные данные
//
//Вывести число без заданных цифр.

package main
import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	var n int
	var copya int = a
	for copya != 0 {
		n++
		copya /= 10
	}
	num := make([]int, n, n)
	for i := 0; i < n; i++ {
		num[n-i-1] = a%10
		a /= 10
	}
	for i := 0; i < len(num); i++ {
		if num[i] == b {
			num = append(num[0:i], num[i+1:]...)
			i--
		}
	}
	for i := 0; i < len(num); i++ {
		fmt.Print(num[i])
	}
}