//Дано трехзначное число. Переверните его, а затем выведите.
//
//Формат входных данных
//На вход дается трехзначное число, не оканчивающееся на ноль.
//
//Формат выходных данных
//Выведите перевернутое число.

package main
import "fmt"
import "math"

func main() {
	var a int
	fmt.Scan(&a)
	var lena int = len(string(a))
	var rev int
	for i := 0; i <= lena; i++ {
		rev += a%10 * int(math.Pow(10, float64(lena - i)))
		a = a / 10
	}
	fmt.Println(rev)
}