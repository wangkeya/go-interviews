package main

import "fmt"

// 模拟动物行为的接口
type IAnimal interface {
	Eat() // 描述吃的行为
}

// 动物 所有动物的父类
type Animal struct {
	Name string
}

// 动物去实现IAnimal中描述的吃的接口
func (a *Animal) Eat() {
	fmt.Printf("%v is eating\n", a.Name)
}

// 动物的构造函数
func NewAnimal(name string) *Animal {
	return &Animal{
		Name: name,
	}
}

// 猫的结构体 组合了animal
type Cat struct {
	*Animal
}

// 实现猫的构造函数 初始化animal结构体
func NewCat(name string) *Cat {
	return &Cat{
		Animal: NewAnimal(name),
	}
}

func main() {
	var cat IAnimal = NewCat("cat")
	cat.Eat() // cat is eating
}
