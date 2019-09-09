package main

// RemoveByIndex will remove element at the specified index
func RemoveByIndex(slice *[]int, index int) {
	*slice = append((*slice)[:index], (*slice)[index+1:]...)
}
