package main

import "fmt"

func whatI() (int, error) {
	fmt.Println("Введите число от 0 до 9")
	var n int
	_, err := fmt.Scanf("%d \n", &n)
	return n, err
}
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
