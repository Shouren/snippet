package main

import (
	"fmt"
)

type T1 struct {
	Name string
}

type T2 struct {
	T1
	Type string
}

type Data struct {
	Name string
	Pass string
}

type TestData struct {
	Datas map[string]Data
}

func testFunc(a Data) {
	a.Name = "Admin"
	a.Pass = "12345"
}

func testFunc0(b map[string]Data) {
	d := b["test"]
	d.Name = "Admin"
	d.Pass = "12345"
	b["test"] = d
}

func main() {
	a := TestData{
		Datas: map[string]Data{},
	}

	a.Datas["test"] = Data{
		Name: "Hello",
		Pass: "World",
	}
	d := a.Datas["test"]
	testFunc(d)

	fmt.Println("a.Datas[\"test\"]")
	fmt.Println(a.Datas["test"].Name)
	fmt.Println(a.Datas["test"].Pass)
	fmt.Println("d")
	fmt.Println(d.Name)
	fmt.Println(d.Pass)

	testFunc0(a.Datas)
	fmt.Println("a.Datas[\"test\"]")
	fmt.Println(a.Datas["test"].Name)
	fmt.Println(a.Datas["test"].Pass)

	fmt.Println("====== Struct initializers test ======")
	t2 := T2{
		T1: T1{
			Name: "test",
		},
		Type: "string",
	}

	fmt.Println(t2.Name)
}
