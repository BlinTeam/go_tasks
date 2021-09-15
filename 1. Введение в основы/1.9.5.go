//Требуется определить, является ли данный год високосным, напомним:
//Год является високосным если он соответствует хотя бы одному из нижеперечисленных условий:
//- кратен 400;
//- кратен 4, но не кратен 100.
//
//Входные данные
//
//Вводится единственное число - номер года (целое, положительное, не превышает 10000).
//
//Выходные данные
//
//Требуется вывести слово YES, если год является високосным и NO - в противном случае.

package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a%400 == 0 || a%4 == 0 && a%100 != 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}