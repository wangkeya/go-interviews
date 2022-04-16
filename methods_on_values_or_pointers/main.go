package main

/*

实例和实例指针可以调用值类型和指针类型 receiver 的方法。
如果通过 method express 方式，struct 值只能调用值类型 receiver 的方法，而 struct 指针是能调用值类型和指针类型 receiver 的方法的。
如果 receiver 是 map、func 或 chan，不要使用指针。
如果 receiver 是 slice，并且方法不会重新分配 slice，不要使用指针。
如果 receiver 是包含 sync.Mutex 或其它类似的同步字段的结构体，receiver 必须是指针，以避免复制。
如果 receiver 是大 struct 或 array，receiver 用指针效率会更高。那么，多大是大？假设要把它的所有元素作为参数传递给方法，如果这样会感觉太大，那对 receiver 来说也就太大了。
如果 receiver 是 struct、array 或 slice，并且它的任何元素都是可能发生改变的内容的指针，最好使用指针类型的 receiver，这会使代码可读性更高。
如果 receiver 是一个本来就是值类型的小 array 或 struct，没有可变字段，没有指针，或只是一个简单的基础类型，如 int 或 string，使用值类型的 receiver 更合适。
值类型的 receiver 可以减少可以生成的垃圾量，如果将值传递给值方法，可以使用栈上的副本而不是在堆上进行分配。编译器会尝试避免这种分配，但不会总成功。不要为此原因却不事先分析而选择值类型的 receiver。
最后，如有疑问，请使用指针类型的 receiver
*/

import (
	"fmt"
)

type Ball struct {
	Name string
}

func (b *Ball) Ping() {
	fmt.Println("ping")
}

func (b Ball) Pong() {
	fmt.Println("pong")
}

func main() {
	v := Ball{}
	p := &Ball{}

	v.Ping()
	v.Pong()

	p.Ping()
	p.Pong()

	//c := Ball{}
	//
	//Ball.Ping(&c)
	//Ball.Pong(c)

	p = &Ball{}

	(*Ball).Ping(p)
	(*Ball).Pong(p)
}
