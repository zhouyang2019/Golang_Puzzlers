package main

import "fmt"

type Cat struct {
	name           string // 名字。
	scientificName string // 学名。
	category       string // 动物学基本分类。
}

func New(name, scientificName, category string) Cat {
	return Cat{
		name:           name,
		scientificName: scientificName,
		category:       category,
	}
}

// 「指针方法」的接收者是该方法所属的那个类型值的「指针值」的一个副本。在该方法内对该副本的修改会体现在原值上
func (cat *Cat) SetName(name string) {
	cat.name = name
}
// 「值方法」的接收者是该方法所属的那个类型值的一个副本。在该方法内对该副本的修改一般不会体现在原值上（除非这个类型本身是某个引用类型）
func (cat Cat) SetName2(name string) {
	cat.name = name
}

func (cat Cat) SetNameOfCopy(name string) {
	cat.name = name
}

func (cat Cat) Name() string {
	return cat.name
}

func (cat Cat) ScientificName() string {
	return cat.scientificName
}

func (cat Cat) Category() string {
	return cat.category
}

func (cat Cat) String() string {
	return fmt.Sprintf("%s (category: %s, name: %q)",
		cat.scientificName, cat.category, cat.name)
}

func main() {
	cat := New("little pig", "American Shorthair", "cat")
	fmt.Printf("The cat: %s\n", cat)
	cat.SetName("monster") // Go语言会自动转译为 (&cat).SetName("monster")
	fmt.Printf("The cat: %s\n", cat)
	cat.SetName2("monster2")
	fmt.Printf("The cat: %s\n", cat)


	cat.SetNameOfCopy("little pig")
	fmt.Printf("The cat: %s\n", cat)

	type Pet interface {
		SetName(name string)
		Name() string
		Category() string
		ScientificName() string
	}

	_, ok := interface{}(cat).(Pet)
	fmt.Printf("Cat implements interface Pet: %v\n", ok)
	_, ok = interface{}(&cat).(Pet)
	fmt.Printf("*Cat implements interface Pet: %v\n", ok)
}
