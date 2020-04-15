package main

import (
	"fmt"

	"github.com/mayone/Go-Utilities/helloworld/buildinfo"
	"github.com/mayone/Go-Utilities/helloworld/config"
	"github.com/mayone/Go-Utilities/helloworld/greet"
)

func main() {
	fmt.Println(buildinfo.Info())
	fmt.Println(config.SayHello())
	fmt.Println(greet.Greet("Sweaty", "Sweety"))
}
