package main

import (
	"crypto/md5"
	"encoding/binary"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"net/http"
	"potatoengine/src/netmessage"

	//"database/sql"
	"fmt"
	"time"

	"potatoengine/src/db"
	message "potatoengine/src/netmessage/pbmessage"
)

func LoginHandle(writer http.ResponseWriter, request *http.Request) {

	fmt.Println("login space is runing")
	var buf, err = ioutil.ReadAll(request.Body)
	if err != nil || buf == nil || len(buf) <= 0 {
		return
	}
	//pb消息
	responsdata := buf[8:]
	var loginRequest = message.LoginResquest{}
	err = proto.Unmarshal(responsdata, &loginRequest)
	if err != nil {
		fmt.Println(err)
		return
	}
	d := buf[4:8]
	id := binary.BigEndian.Uint32(d)
	fn := netmessage.GetProcessFuction(int32(id))
	if fn == nil {
		return
	}
	loginResponse, netError := fn(loginRequest)
	if netError != nil {
		data, _ := netmessage.UnCodePBNetMessage(&netError)
		if data == nil {
			data, _ = netmessage.DefaultNetErrorData()
		}
		writer.Write(data)
		return
	}
	msgbuf, err := netmessage.PackageNetMessage(loginResponse)
	if err != nil {
		return
	}
	//var res,_=loginResponse.(message.LoginResponse)
	writer.Write(msgbuf)
}

///处理登陆
func ProcessLoginRequest(m interface{}) (interface{}, interface{}) {

	var loginRequest, ok = m.(message.LoginResquest)
	if ok == false {
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
	if erro != nil {
		var errinfo = &message.NetError{ErrorCode: message.EMsg_Error_DBClosed}
		return nil, errinfo
	}
	//生成token
	var token string=""
	mkey:=fmt.Sprintf("userid_%d", id)
	token=redis.Get(mkey).Val()
	if token==""{
		to,err:=CenerateToken(name, pass)
		if err != nil {
			var errinfo = &message.NetError{ErrorCode: message.EMsg_Error_DBClosed}
			return nil, errinfo
		}
		token=to
	}
	//把token写入redis
	if err := redis.Set(mkey, token, 0).Err(); err != nil {
		var errinfo = &message.NetError{ErrorCode: message.EMsg_Error_DBClosed}
		return nil, errinfo
	}
	loginResponse := message.LoginResponse{}
	loginResponse.Userid = id
	loginResponse.Token = token
	fmt.Println(token)
	return &loginResponse, nil
}
//根据时间 id password计算toke md5
func CenerateToken(name string, password string) (string, error) {
	if &name == nil || name == "" {
		err := fmt.Errorf("username is nil")
		return "", err
	}
	if password == "" || &password == nil {
		err := fmt.Errorf("password is nil")
		return "", err

	}
	st := fmt.Sprintf("%s%s%s", name, password, time.Now().Unix())
	data := []byte(st)
	mdstr := md5.Sum(data)
	return fmt.Sprintf("%x", mdstr), nil
}
