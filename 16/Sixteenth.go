package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	low, high := 0, len(arr)-1
	point := arr[high]
	point, arr[high] = arr[high], point

	for i := range arr {
		if arr[i] < arr[high] {
			arr[i], arr[low] = arr[low], arr[i]
			low = low + 1
		}
	}
	arr[low], arr[high] = arr[high], arr[low]

	quickSort(arr[:low])
	quickSort(arr[low+1:])
	return arr
}
func main() {
	arr := []int{1, 6, 25, 45, 4, 3, 8, 2, 7, 2}
	fmt.Println(quickSort(arr))
}
