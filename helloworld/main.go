package main

import (
	"fmt"

	"github.com/mayone/Go-Utilities/helloworld/buildinfo"
	"github.com/mayone/Go-Utilities/helloworld/greet"
)

func main() {
	fmt.Println(buildinfo.Info())
	fmt.Println(greet.Greet("Sweaty", "Sweety"))
}
