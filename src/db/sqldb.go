package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

const (
	name     = "root"
	password = "Xiaonan7147!"
	ip       = "106.54.46.165"
	port     = "3306"
	db       = "game"
)

type SQDB struct {
	_sdb  *sql.DB
	_init bool
}

var inst *SQDB

func (db *SQDB) CloseSQL() {
	db._sdb.Close()
}
func (this *SQDB) GetSQL() *sql.DB {
	return this._sdb
}

//func Initialize() {
//
//	//defer inst._sdb.Close() // 延迟关闭 db对象创建成功后才可以调用close方法
//}

func GetSQLManager() *SQDB {
	if inst == nil||inst._init==false {
		inst = &SQDB{
			_sdb: nil,
			_init:true,
		}
		path := strings.Join([]string{name, ":", password, "@tcp(", ip, ":", port, ")/", db, "?charset=utf8"}, "")
		db, err := sql.Open("mysql", path)
		if err != nil {
			println("connect mysql fail")
		}
		inst._sdb = db
		return inst
		//fmt.Println("mysql db is not init")
	}
	return inst
}
