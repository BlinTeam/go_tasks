//Дана строка, содержащая только десятичные цифры. Найти и вывести наибольшую цифру.
//
//Входные данные
//
//Вводится строка ненулевой длины. Известно также, что длина строки не превышает 1000 знаков
//и строка содержит только десятичные цифры.
//
//Выходные данные
//
//Выведите максимальную цифру, которая встречается во введенной строке.

package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	runeNum := rune('9')
	for strings.Contains(text, string(runeNum)) == false {
		runeNum--
	}
	fmt.Println(string(runeNum))
}
