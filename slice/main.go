package main

import "fmt"

func removeByIndex(slice *[]int, index int) {
	*slice = append((*slice)[:index], (*slice)[index+1:]...)
}

func main() {
	sliceInMain := []int{5, 2, 3, 1, 4}
	fmt.Println(sliceInMain)
	removeByIndex(&sliceInMain, 1)
	fmt.Println(sliceInMain)
}
