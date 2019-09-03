package swap

// Swap two integers
func Swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}
