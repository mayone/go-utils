package main

import "fmt"

// RemoveByIndex will remove element at the specified index
func RemoveByIndex(slice *[]int, index int) {
	*slice = append((*slice)[:index], (*slice)[index+1:]...)
}

func main() {
	sliceInMain := []int{5, 2, 3, 1, 4}
	fmt.Println(sliceInMain)
	RemoveByIndex(&sliceInMain, 1)
	fmt.Println(sliceInMain)
}
