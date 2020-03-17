package main

import (
	"fmt"

	"github.com/mayone/Go-Utilities/helloworld/build"
	"github.com/mayone/Go-Utilities/helloworld/greet"
)

func main() {
	fmt.Println(build.Info())
	fmt.Println(greet.Greet("Sweaty", "Sweety"))
}
