package main

import (
	"fmt"
)

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
	b["test"].Name = "Admin"
	b["test"].Pass = "12345"
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
}
