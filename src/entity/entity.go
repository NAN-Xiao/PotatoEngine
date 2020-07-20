package entity

import (
	"fmt"
	"potatoengine/src/logService"
	"potatoengine/src/netWork"
	"potatoengine/src/space"
)

type Entity struct {
	Conn          netWork.IConn
	EntityID      int32
	spaceid       int32
	reveiveMsgQue chan interface{}
	sendMsgQue    chan interface{}
	Created       bool
}
//开始接受发送线程
func (this *Entity) Connect() {
	if this.Conn == nil {
		return
	}
	//接收消息放入 receive队列
	go this.Conn.Receive(this.reveiveMsgQue)
	//发送消息到队列 从send队列
	go func(msgque chan interface{}, con netWork.IConn) {
		for {
			if msgque == nil {
				break
			}
			if len(msgque) <= 0 {
				continue
			}
			msg := <-msgque
			con.Send(msg)
		}
	}(this.sendMsgQue, this.Conn)
}

//从队列读取消息
func (this *Entity) Read() interface{} {
	if this.reveiveMsgQue == nil || len(this.reveiveMsgQue) <= 0 {
		return nil
	}
	return <-this.reveiveMsgQue
}

//放入消息队列pkg
func (this *Entity) Write(pkg interface{}) {

	if this.Created == false {
		logService.LogError(fmt.Sprintf("this entity is not init,this.conn id is ::%d", this.Conn.GetID()))
	}
	if pkg != nil {
		this.sendMsgQue <- pkg
	}
}

func (this *Entity) GetEntityID() int32 {
	return this.EntityID
}
func (this *Entity) SetEntityID(id int32) {
	this.EntityID = id
}
func (this *Entity) GetSpaceID() int32 {
	return this.spaceid
}

//得到当前所在是space
func (this *Entity) GetCurrentSpace() space.ISpace {
	sp := space.GetSpaceByID(this.spaceid)
	if sp == nil {
		return nil
	}
	return sp
}

//进入指定场景
func (this *Entity) EnterSpace(sp space.ISpace) {

	nspace := space.GetSpaceByID(sp.GetSpace().SpaceID)
	if nspace == nil {
		logService.LogError(fmt.Sprintf("this entity ready to enter next space is nil ,this.conn id is ::%s", this.Conn.GetID()))
	}
	nspace.GetSpace().EnterSpace(this)
}

//退出指定场景
func (this *Entity) LeaveSpace(sp space.ISpace) {

	cspace := space.GetSpaceByID(sp.GetSpace().SpaceID)
	if cspace == nil {
		logService.LogError(fmt.Sprintf("this entity is not at any space,this.conn id is ::%s", this.Conn.GetID()))
	}
	cspace.GetSpace().LeaveSpace(this)
}

//创建一个新到entity
func (this *Entity) CreatEntity(conn netWork.IConn) {
	this.EntityID = -1
	this.Conn = conn
	this.spaceid = -1
	this.reveiveMsgQue = make(chan interface{}, 128)
	this.sendMsgQue = make(chan interface{}, 128)
	this.Created = true
}
