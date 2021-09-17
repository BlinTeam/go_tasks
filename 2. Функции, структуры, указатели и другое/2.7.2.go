//Дана строка, содержащая только английские буквы (большие и маленькие). Добавить символ ‘*’ (звездочка) между буквами
//(перед первой буквой и после последней символ ‘*’ добавлять не нужно).
//
//Входные данные
//
//Вводится строка ненулевой длины. Известно также, что длина строки не превышает 1000 знаков.
//
//Выходные данные
//
//Вывести строку, которая получится после добавления символов '*'.

package main
import (
	"fmt"
	"os"
	"bufio"
	"unicode/utf8"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	bs := []rune(text)
	for i := 0; i < utf8.RuneCountInString(text) - 1; i++ {
		fmt.Print(string(bs[i]), "*")
	}
	fmt.Print(string(bs[utf8.RuneCountInString(text) - 1]))
}