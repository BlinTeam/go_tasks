//Поменяйте местами значения переменных на которые ссылаются указатели. После этого переменные нужно вывести.
//ВАЖНО: Считайте что пакет main уже объявлен, а также функция main() уже есть.

func test(x1 *int, x2 *int) {
	c := *x1
	*x1 = *x2
	*x2 = c
	fmt.Print(*x1, *x2)
}