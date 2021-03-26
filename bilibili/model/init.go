package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB
var err error

func Initdb() *gorm.DB {
	DB, err = gorm.Open("mysql", "jingzhe:123@/bilibili?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	return DB
}
