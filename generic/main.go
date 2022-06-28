package main

import "fmt"

func Print[T any](s ...T) {
	for _, v := range s {
		fmt.Print(v)
	}
}

func Equal[E comparable](s1, s2 []E) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, v1 := range s1 {
		v2 := s2[i]
		if v1 != v2 {
			return false
		}
	}
	return true
}

func Clone[S ~[]E, E any](s S) S {
	// Preserve nil in case it matters.
	if s == nil {
		return nil
	}
	return append(S([]E{}), s...)
}

func main() {
	Print("Hello, ", "playground\n")
	Print(Equal([]string{"1"}, []string{"1"}))
	Print("\n")

	Print(Clone([]string{"222"}))
}
