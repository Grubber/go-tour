package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Print("OS X.")
	case "linux":
		fmt.Print("Linux.")
	default:
		fmt.Printf("%s.", os)
	}
}
