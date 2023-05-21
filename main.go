package main

import (
	"fmt"
	"runtime"
)

var version = "0.0.0-dev"

func main() {
	fmt.Println("hello world")
	fmt.Println(runtime.GOOS, runtime.GOARCH)
	fmt.Println("version: ", version)
}
