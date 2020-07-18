package entity

import (
	"fmt"
	"potatoengine/src/logService"
	"potatoengine/src/space"
)

type Entity struct {
	Created       bool
	connID        int32
	Agnetid       int32
	spaceid       int32
	reveiveMsgQue chan interface{}
	sendMsgQue    chan interface{}
}

//消息收发
func (this *Entity) Read() interface{} {
	if this.reveiveMsgQue == nil || len(this.reveiveMsgQue) <= 0 {
		return nil
	}
	return <-this.reveiveMsgQue
}

func (this *Entity) Write(pkg interface{}) {

	if this.Created == false {
		logService.LogError(fmt.Sprintf("this entity is not init,this.conn id is ::%s", this.connID))
	}
	if pkg != nil {
		this.sendMsgQue <- pkg
	}
}

//
func (this *Entity) GetEntityID() int32 {
	return this.id
}
func (this *Entity) SetEntityID(id int32) {
	this.id = id
}
func (this *Entity) GetSpaceID() int32 {
	return this.spaceid
}

//当前所在是space
func (this *Entity) GetCurrentSpace() *space.BaseSpace {
	sp := space.GetSpace(this.spaceid)
	if sp == nil {
		return nil
	}
	return sp
}

//进入场景
func (this *Entity) EnterSpace(spaceID int32) {

	nspace := space.GetSpace(spaceID)
	if nspace == nil {
		logService.LogError(fmt.Sprintf("this entity ready to enter next space is nil ,this.conn id is ::%s", this.connID))
	}
	nspace.EnterSpace(this)
}

//退出场景
func (this *Entity) LeaveSpace(spaceID int32) {

	cspace := space.GetSpace(this.spaceid)
	if cspace == nil {
		logService.LogError(fmt.Sprintf("this entity is not at any space,this.conn id is ::%s", this.connID))
	}
	cspace.LeaveSpace(this)
}
