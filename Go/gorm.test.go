package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type TestData struct {
	gorm.Model
	Name string
	Desc string
}

func main() {
	user := "root"
	passwd := "0192837465"
	dbname := "test"
	config := user + ":" + passwd + "@tcp(localhost)/" + dbname
	config += "?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open("mysql", config)
	if err != nil {
		panic("Failed to connect to DB: " + err.Error())
	}
	defer db.Close()

	db.AutoMigrate(&TestData{})

	db.Create(&TestData{
		Name: "One",
		Desc: "Test data for gorm, test",
	})

	var data TestData
	db.Where("name = ?", "One").First(&data)
	fmt.Println(data)

	// db.Delete(&data)
}
