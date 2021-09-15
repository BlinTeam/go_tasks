//По данному трехзначному числу определите, все ли его цифры различны.
//
//Формат входных данных
//На вход подается одно натуральное трехзначное число.
//
//Формат выходных данных
//Выведите "YES", если все цифры числа различны, в противном случае - "NO".

package main

import "fmt"

func main(){

	var a int
	fmt.Scan(&a)

	var b int = a % 10
	var c int = a % 100 / 10
	var d int = a / 100
	var res bool = false

	if b != c {
		if b != d {
			if d != c {
				res = true
			}
		}
	}

	if res == true{
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}