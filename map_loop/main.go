package main

import "fmt"

// go tool compile -S main.go
func main() {
	ageMp := make(map[string]int)
	ageMp["qcrao"] = 18
	ageMp["qhi"] = 20
	ageMp["qbac"] = 90
	ageMp["qbae"] = 90
	ageMp["qbad"] = 90
	ageMp["qbag"] = 90

	for name, age := range ageMp {
		fmt.Println(name, age)
	}
}
