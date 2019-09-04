package main

import (
	"fmt"
	"time"
)

func readword(ch chan string) {
	fmt.Println("Type a word, then hit Enter.")
	var word string
	fmt.Scanf("%s", &word)
	ch <- word
}

func timeout(t chan bool) {
	time.Sleep(5 * time.Second)
	t <- false
}

func main() {
	a := 3
	b := 4
	c := 5
	d := []int{a, b, c}
	for a < 2 || b > 3 {
		fmt.Println("Hello")
	}
}
