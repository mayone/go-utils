package slice

// removeByIndex will remove element at the specified index
func removeByIndex(slice *[]int, index int) {
	*slice = append((*slice)[:index], (*slice)[index+1:]...)
}
