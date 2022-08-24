package config

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	d, err := gorm.Open("mysql", "sqluser:mysql/goapi?charset=utf&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	db = d
}

func GetDb() *gorm.DB {
	return db
}