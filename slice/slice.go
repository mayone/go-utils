package slice

// removeByIndex will return slice with removal of element at the specified index
func removeByIndex(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

// removeByIndex will remove element at the specified index
// func removeByIndex(slice *[]int, index int) {
// 	*slice = append((*slice)[:index], (*slice)[index+1:]...)
// }
