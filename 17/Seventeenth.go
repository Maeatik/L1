package main

import "fmt"

func binarySearch(target int, arr []int) bool {
	low, high := 0, len(arr)-1
	for low <= high {
		middle := (low + high) / 2

		if arr[middle] >= target {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}
	if low == len(arr) || arr[low] != target {
		return false
	}
	return true
}

func main() {
	items := []int{1, 2, 3, 4, 6, 7, 8, 25, 45}
	fmt.Println(binarySearch(2, items))
}
