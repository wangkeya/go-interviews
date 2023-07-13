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
	demo := example1("demo")
	fmt.Println(demo)

	example2()

	in := example3()
	fmt.Println(in()) // 1
	fmt.Println(in()) // 2
}

type Demo struct {
	name string
}

/*
*
1、指针逃逸应该是最容易理解的一种情况了，即在函数中创建了一个对象，返回了这个对象的指针。
2、空接口即 interface{} 可以表示任意的类型，如果函数参数为 interface{}，编译期间很难确定其参数的具体类型，也会发生逃逸
*/
func example1(name string) *Demo {
	d := new(Demo) // 局部变量 d 逃逸到堆
	d.name = name
	return d
}

/*
*
3、操作系统对内核线程使用的栈空间是有大小限制的，64 位系统上通常是 8 MB。可以使用 ulimit -a 命令查看机器上栈允许占用的内存的大小
*/
func example2() {
	generate8191()
	generate8192()
	generate(1)
}

func generate8191() {
	nums := make([]int, 8191) // < 64KB
	for i := 0; i < 8191; i++ {
		nums[i] = rand.Int()
	}
}

func generate8192() {
	nums := make([]int, 8192) // = 64KB
	for i := 0; i < 8192; i++ {
		nums[i] = rand.Int()
	}
}

func generate(n int) {
	nums := make([]int, n) // 不确定大小
	for i := 0; i < n; i++ {
		nums[i] = rand.Int()
	}
}

/*
*
4、闭包, 让你可以在一个内层函数中访问到其外层函数的作用域
*/
func example3() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}
