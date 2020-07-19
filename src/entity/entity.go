package entity

import (
	"fmt"
	"potatoengine/src/connection"
	"potatoengine/src/logService"
	"potatoengine/src/space"
)

type Entity struct {
	connID        connection.ConnID
	EntityID      int32
	spaceid       int32
	reveiveMsgQue chan interface{}
	sendMsgQue    chan interface{}

	Created bool
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
	return this.EntityID
}
func (this *Entity) SetEntityID(id int32) {
	this.EntityID = id
}
func (this *Entity) GetSpaceID() int32 {
	return this.spaceid
}

//当前所在是space
func (this *Entity) GetCurrentSpace() *space.BaseSpace {
	sp := spaceMgr.GetSpace(this.spaceid)
	if sp == nil {
		return nil
	}
	return sp
}

//进入场景
func (this *Entity) EnterSpace(spaceID int32) {

	nspace := spaceMgr.GetSpace(spaceID)
	if nspace == nil {
		logService.LogError(fmt.Sprintf("this entity ready to enter next space is nil ,this.conn id is ::%s", this.connID))
	}
	nspace.EnterSpace(this)
}

//退出场景
func (this *Entity) LeaveSpace(spaceID int32) {

	cspace := spaceMgr.GetSpace(this.spaceid)
	if cspace == nil {
		logService.LogError(fmt.Sprintf("this entity is not at any space,this.conn id is ::%s", this.connID))
	}
	cspace.LeaveSpace(this)
}

func (this *Entity) CreatEntity(connid connection.ConnID) {
	this.connID = connid
	this.EntityID=-1
	this.spaceid = -1
	this.reveiveMsgQue=make(chan interface{},128)
	this.sendMsgQue=make(chan interface{},128)
	this.Created=true

}
