package main

import (
	"fmt"
)

func main() {
	A := make(chan bool, 1)
	B := make(chan bool)

	Exit := make(chan bool)

	go func() {
		for i := 0; i < 26; i++ {
			if ok := <-A; ok {
				fmt.Println(string(rune('a' + i)))
				B <- true
			}
		}
	}()

	go func() {
		defer func() {
			close(Exit)
		}()
		for i := 0; i < 26; i++ {
			if ok := <-B; ok {
				fmt.Println(string(rune('A' + i)))
				A <- true
			}
		}
	}()

	A <- true
	<-Exit
}
