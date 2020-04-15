// +build dev

package config

import "fmt"

const (
	ENV = "DEV"
)

func SayHello() string {
	return fmt.Sprintf("Hello, this is a %s build", ENV)
}
