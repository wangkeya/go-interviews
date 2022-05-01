package main

import (
	"fmt"
	"math/rand"
)

// go build -gcflags="-m" main.go
/*
# command-line-arguments
./main.go:17:18: inlining call to rand.Intn
./main.go:12:13: inlining call to fmt.Println
./main.go:17:2: moved to heap: tmp
./main.go:12:13: ... argument does not escape
./main.go:12:14: *num escapes to heap

*/

func main() {
	num := GenerateRandomNum()
	fmt.Println(*num)
}

//go:noinline
func GenerateRandomNum() *int {
	tmp := rand.Intn(500)

	return &tmp
}
