//Напишите функцию, находящую наименьшее из четырех введённых в этой же функции чисел.
//
//Входные данные
//
//Вводится четыре числа.
//
//Выходные данные
//
//Необходимо вернуть из функции наименьшее из 4-х данных чисел

func minimumFromFour() int {
	var arr [4]int

	for i:= 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}
	var min int = arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}