package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"net/http"
	"potatoengine/src/agent"
	"potatoengine/src/db"
	"potatoengine/src/message"
	"potatoengine/src/space"
)

type Handle interface {
	//ServeHTTP(ResponseWriter, *Request)
}

type GateSpace struct {
	space.BaseSpace
}

//todo http监听返回登陆结果
func (this *GateSpace) Process() {
	http.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {

		var matching = false
		go func() {
			buffer, err := ioutil.ReadAll(request.Body)
			if err != nil {
				return
			}

			er := proto.Unmarshal(buffer, nil)
			if er != nil {
				return
			}
			sql := db.GetSQLManager().GetSQL()
			if sql == nil {
				return
			}
			var username string
			var userpassword string

			rows, err := sql.Query(" SELECT * FROM user")
			for rows.Next() {
				var id int
				var name string
				var password string
				err = rows.Scan(&id, &name, &password)
				if username == name && password == userpassword {
					matching = true
					break
				}
				fmt.Println(id, name, password)
			}

			if matching == true {
				writer.Write([]byte{1})
				return
			}
			writer.Write([]byte{0})
		}()
	})
	http.ListenAndServe("0.0.0.0:8999", nil)
}

//agent 进入场景
func (this *GateSpace) LeaveSpace(ag *agent.Agent) {
	v, ok := this.Agents[ag.Aid]
	if ok {
		v.OnLeaveSpace()
		delete(this.Agents, v.Aid)
	}
}

//agent退出场景
func (this *GateSpace) EnterSpace(ag *agent.Agent) {
	_, ok := this.Agents[ag.Aid]
	if ok {
		return
	}
	this.Agents[ag.Aid] = ag
	ag.OnLeaveSpace()
}

func (this *GateSpace) GetName() string {
	return this.Spacename
}

func NewGateSpace(name string) space.ISpace {
	sp := &LoginSpace{space.BaseSpace{
		SpaceID:    0,
		Spacename:  name,
		Agents:     make(map[uint32]*agent.Agent),
		Spacechanl: make(chan *message.MsgPackage, 100),
	}}
	return sp
}
