//Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования.
//Длина пароля должна быть не менее 5 символов, он должен содержать только цифры и буквы латинского алфавита.
//На вход подается строка-пароль. Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"

package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
	"os"
	"bufio"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	rs := []rune(text)

	if utf8.RuneCountInString(text) < 5 {
		fmt.Println("Wrong password")
		return
	}
	for i := 0; i < utf8.RuneCountInString(text); i++ {
		if unicode.IsDigit(rs[i]) == false && unicode.Is(unicode.Latin, rs[i]) == false {
			fmt.Println("Wrong password")
			return
		}
	}
	fmt.Print("Ok")
}