package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"potatoengine/src/db"
	"potatoengine/src/netmessage"
	message "potatoengine/src/netmessage/pbmessage"
)

func ProcessLoginRequest(m interface{}) ( interface{}, interface{}) {


	var loginRequest,ok=m.(message.LoginResquest)
	if ok==false{
		var errinfo = &message.NetError{ErrorCode: message.EMsg_Error_UserInfo}
		return nil, errinfo
	}

	var id int32
	var name string
	var pass string
	sql := db.GetSQLManager().GetSQL()
	if sql == nil || sql.Ping() != nil {
		//发送服务器错误消息
		var errinfo = &message.NetError{ErrorCode: message.EMsg_Error_DBClosed}
		return nil, errinfo
	}

	myquery := fmt.Sprintf("SELECT * FROM `game`.`userinfo` WHERE `username` = '%s'AND`password`='%s'", loginRequest.Username, loginRequest.Password)
	rows := sql.QueryRow(myquery)
	//mysql错误
	if rows != nil {
		rows.Scan(&id, &name, &pass)
		if id <= 0 {
			var errinfo = &message.NetError{ErrorCode: message.EMsg_Error_DBClosed}
			return nil, errinfo
		}
	}
	//生成token返回token
	//存到redis key是id value是token
	redis, erro := db.GetRedisManager().GetDB()
	//检查redis
	if erro == nil {
		var errinfo = &message.NetError{ErrorCode: message.EMsg_Error_DBClosed}
		return nil, errinfo
	}
	//生成token
	var token string
	token, err := CenerateToken(name, pass)
	if err == nil {
		var errinfo = &message.NetError{ErrorCode: message.EMsg_Error_DBClosed}
		return nil, errinfo
	}
	//把token写入redis
	if err := redis.Set(fmt.Sprintf("userid_%d", id), token, 0).Err(); err != nil {
		var errinfo = &message.NetError{ErrorCode: message.EMsg_Error_DBClosed}
		return nil, errinfo
	}
	loginResponse:=message.LoginResponse{}
	loginResponse.Userid=id
	loginResponse.Token=token

}
