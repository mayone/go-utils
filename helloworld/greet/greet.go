package greet

import "fmt"

func Greet(from, to string) string {
	return fmt.Sprintf("Hello %s, I am %s", to, from)
}
