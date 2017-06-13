package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Hello World pid=%d\n", os.Getpid())

	count := 10
	for i := 0; i < count; i++ {
		fmt.Printf("%d\n", i)
	}
}
