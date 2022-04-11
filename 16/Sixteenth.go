package main
//Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
import "fmt"

func quickSort(arr []int) []int {
	//если длина меньше 2 - 0 или 1, просто возращается массив
	if len(arr) < 2 {
		return arr
	}
	//берутся первое и последнее число
	low, high := 0, len(arr)-1

	for i := range arr {
		//если число меньше крайнего правого они меняются местами, и указатель на левое
		//необработанное число двигается вправо
		if arr[i] < arr[high] {
			arr[i], arr[low] = arr[low], arr[i]
			low = low + 1
		}
	}
	//необработанное левое и правое число - меняются
	arr[low], arr[high] = arr[high], arr[low]
	//массив разделяется на две части, до и после необработанного левого числа
	quickSort(arr[:low])
	quickSort(arr[low+1:])

	return arr
}
func main() {
	arr := []int{1, 6, 25, 45, 4, 3, 8, 2, 7, 2}
	fmt.Println(quickSort(arr))
}
