package gorrm

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Conn() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/waf?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("数据库连接成功")
	}

	DB = db
}
