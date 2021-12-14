package main

import "fmt"

// 接口类型与其他数据类型不同，它是没法实例化的
// 对于某一个接口来说，如果没有任何数据类型可以作为它的实现，那么该接口的值就不可能存在
type Pet interface {
	// 接口类型花括号包裹的是它的「方法定义」，一个接口的方法集合就是它的全部特征
	SetName(name string)
	Name() string
	Category() string
}

type Dog struct {
	// 结构体类型包裹的是它的「字段声明」
	name string // 名字。
}

// Duck typing，对于任何数据类型，只要它的方法集合中完全包含了一个接口的全部特征（即全部的方法），那么它一定是这个接口的实现类型
// 方法的签名完全一致；方法的名称要一摸一样
func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func (dog Dog) Category() string {
	return "dog"
}

type Cat struct {
	name string
	sex string
}
func (cat Cat) SetName(name string) {
	cat.name = name
}
func (cat Cat) Category() string {
	return "cat"
}
func (cat Cat) Name() string {
	return cat.name
}
func (cat Cat) Name2() string {
	return cat.name
}

func main() {
	// 示例1。
	dog := Dog{"little pig"}
	_, ok := interface{}(dog).(Pet)
	fmt.Printf("Dog implements interface Pet: %v\n", ok)
	_, ok = interface{}(&dog).(Pet)
	fmt.Printf("*Dog implements interface Pet: %v\n", ok)
	fmt.Println()

	// 示例2。
	var pet Pet = &dog
	fmt.Printf("This pet is a %s, the name is %q.\n",
		pet.Category(), pet.Name())
	fmt.Println()

	cat := Cat{"little cat", "female"}
	_, ok = interface{}(cat).(Pet)
	fmt.Printf("Cat implements interface Pet: %v\n", ok)
	_, ok = interface{}(&cat).(Pet)
	fmt.Printf("*Cat implements interface Pet: %v\n", ok)
	fmt.Println()

}
