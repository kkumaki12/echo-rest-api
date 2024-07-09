package main

import (
	"echo-rest-api/db"
	"echo-rest-api/model"
	"fmt"
)

func main() {
	dbConn := db.NewDb() // DBのインスタンスのアドレスが返却される
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDb(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{}) // 反映させたいモデル構造を記載
}
