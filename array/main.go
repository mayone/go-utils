package main

import "fmt"

func removeByIndex(arr *[]int, index int) {
	*arr = append((*arr)[:index], (*arr)[index+1:]...)
}

func main() {
	arr := []int{5, 2, 3, 1, 4}
	fmt.Println(arr)
	removeByIndex(&arr, 1)
	fmt.Println(arr)
}
