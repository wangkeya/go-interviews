package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
)

func main() {
	/*
		内置的len函数可以返回一个字符串中的字节数目(不是rune字符数目)
		rune是utf8编码的每一个的字符

		rune能处理一切的字符，而byte仅仅局限在ascii
	*/
	a := "aaa我的"

	fmt.Println(len(a))
	fmt.Println(utf8.RuneCountInString(a))

	reflect.TypeOf(a)

	fmt.Println(reflect.TypeOf(a))
}
