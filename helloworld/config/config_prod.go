// +build !dev

package config

import "fmt"

const (
	ENV = "PROD"
)

func SayHello() string {
	return fmt.Sprintf("Hello, this is a %s build", ENV)
}
