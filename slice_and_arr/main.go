package main

import "fmt"

/*
区别：
1、切片是指针类型，数组是值类型；
2、数组的长度是固定的，而切片不是（切片可以看成动态的数组）；
3、切片比数组多一个容量（cap）属性；
4、切片的底层是数组
*/
func main() {
	a := [2]int{8, 6}
	b := [2]int{5, 6}
	if a == b {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}

	/*
	   if a[:] == b[:] {
	       fmt.Println("equal")
	   } else {
	       fmt.Println("not equal")
	   }
	*/

}
