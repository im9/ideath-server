package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Users ユーザー情報のテーブル情報
type Users struct {
    ID       int
    Name     string `json:"name"`
    Age      string `json:"age"`
    Address  string `json:"address"`
    UpdateAt string `json:"updateAt" sql:"not null;type:date"`
}

func main() {
	_, err := sqlConnect()
	if err != nil {
			panic(err.Error())
	} else {
			fmt.Println("DB接続成功")
	}
}

func sqlConnect() (database *gorm.DB, err error) {
	DBMS := "mysql"
	USER := "root"
	PASS := ""
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "mypage_test"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	return gorm.Open(DBMS, CONNECT)
}
