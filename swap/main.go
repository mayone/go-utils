package main

import (
	"fmt"
)

func Swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func main() {
	a := 3
	b := 5
	fmt.Println("Before swapping")
	fmt.Println("a =", a, "b =", b)
	Swap(&a, &b)
	fmt.Println("After swapping")
	fmt.Println("a =", a, "b =", b)
}
