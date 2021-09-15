//В рамках этого урока мы постарались представить себе уже привычные нам переменные и функции, как объекты из реальной жизни.
//Чтобы закрепить результат мы предлагаем вам небольшую творческую задачу.
//
//Вам необходимо реализовать структуру со свойствами-полями On, Ammo и Power, с типами bool, int, int соответственно.
//У этой структуры должны быть методы: Shoot и RideBike, которые не принимают аргументов, но возвращают значение bool.
//
//Если значение On == false, то оба метода вернут false.
//
//Делать Shoot можно только при наличии Ammo (тогда Ammo уменьшается на единицу, а метод возвращает true), если его нет,
//то метод вернет false. Метод RideBike работает также, но только зависит от свойства Power.
//
//Чтобы проверить, что вы все сделали правильно, вы должны создать указатель на экземпляр этой структуры с именем testStruct в функции main,
//в дальнейшем программа проверит результат.
//
//Закрывающая фигурная скобка в конце main() вам не видна, но она есть.
//
//Пакет main объявлять не нужно!

type abc struct {
	On bool
	Ammo int
	Power int
}

func (str *abc) Shoot() bool {
	if str.On == false {
		return false
	}
	if str.Ammo == 0 {
		return false
	}
	str.Ammo --
	return true
}

func (str1 *abc) RideBike() bool {
	if str1.On == false {
		return false
	}
	if str1.Power == 0 {
		return false
	}
	str1.Power --
	return true
}

func main() {
	testStruct := new(abc)