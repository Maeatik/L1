package main

//Удалить i-ый элемент из слайса.

import "fmt"

func whatI() (int, error) {
	fmt.Println("Введите число от 0 до 9")
	var n int
	_, err := fmt.Scanf("%d \n", &n)
	return n, err
}

//быстрый вариант - он меняет выбранный для удаления элемент и крайний правый.
//После этого, крайний правый элемент удаляется
func vers_1(slice []string) {
	i, err := whatI()
	if err != nil {
		return
	}
	slice[i] = slice[len(slice)-1]
	slice[len(slice)-1] = ""
	slice = slice[:len(slice)-1]
	fmt.Println(slice)
}

//медленный вариант. В слайс копируется этот же слайс, но без i-го элемента. Чем больше слайс - тем дольше выполняется,
//но зато последний элемент слайса не в середине
func vers_2(slice []string) {
	i, err := whatI()
	if err != nil {
		return
	}
	copy(slice[i:], slice[i+1:])
	slice[len(slice)-1] = ""
	slice = slice[:len(slice)-1]
	fmt.Println(slice)
}

//К слайсу добвляется его вторая половина, кроме i-го числа
func vers_3(slice []string) {
	i, err := whatI()
	if err != nil {
		return
	}
	slice = append(slice[:i], slice[i+1:]...)

	fmt.Println(slice)
}
func main() {
	slice_1 := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	vers_1(slice_1)

	slice_2 := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	vers_2(slice_2)

	slice_3 := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	vers_3(slice_3)
}
