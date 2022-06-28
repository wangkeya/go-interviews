package main

import "fmt"

func Print[T interface{}](s ...T) {
	for _, v := range s {
		fmt.Print(v)
	}
}

func main() {
	Print("Hello, ", "playground\n")
}
