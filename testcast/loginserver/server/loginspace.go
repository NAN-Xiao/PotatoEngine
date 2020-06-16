package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"potatoengine/src/agent"
	"potatoengine/src/db"
	"potatoengine/src/message"
	"potatoengine/src/space"
)

//Json字段必须大写
type UserInfo struct {
	Name string `json:"username"`
	Pass string `json:"password"`
}

type LoginSpace struct {
	space.BaseSpace
}

func Login(writer http.ResponseWriter, request *http.Request) {
	// print(request)
	buf, err := ioutil.ReadAll(request.Body)
	fmt.Println("Server response")
	//panic(err)
	if err != nil || buf == nil || len(buf) <= 0 {
		print(err)
		return
	}
	//fmt.Println(string(buf))
	var info UserInfo
	err = json.Unmarshal(buf, info)
	fmt.Println("222", info.Name)

	if err != nil {
		print("Unmarshal err")
		return
	}
	sql := db.GetSQLManager().GetSQL()
	if sql == nil {
		return
	}
	myquery := fmt.Sprintf("SELECT * FROM user Where 'username'=='%s' AND 'password'=='%s'", info.Name, info.Pass)
	row := sql.QueryRow(myquery)

	var name string
	var password string
	err = row.Scan(&name, &password)
	if err != nil {

	}
	print(name)
	print(password)
	// for rows.Next() {
	// 	var id int
	// 	var name string
	// 	var password string
	// 	err = rows.Scan(&id, &name, &password)
	// 	if username == name && password == userpassword {
	// 		matching = true
	// 		break
	// 	}
	// 	fmt.Println(id, name, password)
	// }

	// if matching == true {
	// 	writer.Write([]byte{1})
	// 	return
	// }
	// writer.Write([]byte{0})
}

//todo http监听返回登陆结果
func (this *LoginSpace) Process() {
	http.HandleFunc("/login", Login)
	http.ListenAndServe("0.0.0.0:8999", nil)
	//fmt.Println("loging server started")
}

//agent 进入场景
func (this *LoginSpace) LeaveSpace(ag *agent.Agent) {
	v, ok := this.Agents[ag.GetUserID()]
	if ok {
		v.OnLeaveSpace()
		delete(this.Agents, v.GetUserID())
	}
}

//agent退出场景
func (this *LoginSpace) EnterSpace(ag *agent.Agent) {
	_, ok := this.Agents[ag.GetUserID()]
	if ok {
		return
	}
	this.Agents[ag.GetUserID()] = ag
	ag.OnLeaveSpace()
}

func (this *LoginSpace) GetName() string {
	return this.Spacename
}

func NewLoginSpace(name string) space.ISpace {
	sp := &LoginSpace{space.BaseSpace{
		SpaceID:    0,
		Spacename:  name,
		Agents:     make(map[uint32]*agent.Agent),
		Spacechanl: make(chan *message.MsgPackage, 100),
	}}
	return sp
}
