package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type SQDB struct {
	_sdb *sql.DB
}

var inst *SQDB

func (db *SQDB) CloseSQL() {
	db._sdb.Close()
}


func Initialize() {
	inst := &SQDB{
		_sdb: nil,
	}
	db, err := sql.Open("mysql", "root:xxxxxx@tcp(118.24.159.133:3306)/student?charset=utf8")
	if err != nil {
		println("connect mysql fail")
	}
	inst._sdb = db
	//defer inst._sdb.Close() // 延迟关闭 db对象创建成功后才可以调用close方法
}

func GetSQLManager() *SQDB {
	if inst == nil {
		fmt.Println("mysql db is not init")
		return nil
	}
	return inst
}
