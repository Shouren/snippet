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

func (t *TestData) GetByName(db *gorm.DB, name string) error {
	t.Name = name
	err := db.First(t).Error
	if err != nil {
		return err
	}
	return nil
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

	// db.Create(&TestData{
	// 	Name: "One",
	// 	Desc: "Test data for gorm, test",
	// })

	var data TestData
	err = data.GetByName(db, "One")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(data)
	}

	// db.Delete(&data)
}
