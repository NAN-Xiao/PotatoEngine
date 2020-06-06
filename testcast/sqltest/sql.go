package main

import (
	"fmt"
	"potatoengine/src/db"
)

//测试sqlmagr
func main() {

	sql := db.GetSQLManager()
	if sql == nil {
		fmt.Println("de not open")
		return
	}
	sdb := sql.GetSQL()
	err := sdb.Ping()
	if err != nil {
		fmt.Println("not connect sql")
		return
	}
	sql2 := db.GetSQLManager()
	sdb2 := sql2.GetSQL()
	err2 := sdb2.Ping()
	if err2 != nil {
		return
	}

	rows, err := sdb2.Query(" SELECT * FROM user")
	for rows.Next() {
		var id int
		var name string
		var password string
		err = rows.Scan(&id, &name, &password)
		fmt.Println(id, name, password)
	}
	fmt.Println("db opend")

}
