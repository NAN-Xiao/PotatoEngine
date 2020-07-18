package space

import (
	"fmt"
	"potatoengine/src/entity"
	"potatoengine/src/logService"
	"potatoengine/src/netmessage"
	spaceMgr "potatoengine/src/space"
)

type BaseSpace struct {
	GameID int32
	SpaceID    int32
	Spacename  string
	Entitys    map[int32]*entity.Entity
	Spacechanl chan netmessage.ServerMsgPackage
}

//从当前space移除entity
func (this *BaseSpace) LeaveSpace(et *entity.Entity) {
	if len(this.Entitys) <= 0 || this.Entitys == nil {
		return
	}
	id := et.GetEntityID()
	_, ok := this.Entitys[id]
	if ok {
		delete(this.Entitys, id)
	}
}
//entity加入当前地图
func (this *BaseSpace) EnterSpace(entity *entity.Entity) {
	if len(this.Entitys) <= 0 || this.Entitys == nil {
		logService.LogError(fmt.Sprintf("space(id:%s) is not have entity map", this.SpaceID))
		return
	}
	id := entity.GetEntityID()
	if sp := spaceMgr.GetSpace(id); sp==nil {
		logService.LogError(fmt.Sprintf("space(id:%s) have same entity(entityid :%s)", this.SpaceID, id))
		return
	}
	this.Entitys[id] = entity
}

//暂时不动
func (this *BaseSpace)SetGameID(gid int32){

}
func (this *BaseSpace)GetGameID()int32{

}
func (this *BaseSpace)GetName() string{

}
func (this *BaseSpace)GetID() int32{

}
//开始启动调用
func (this *BaseSpace)OnStart(){

}
//按时间间隔调用
func (this *BaseSpace)Tick(){

}
//不按时间同步调用
func (this *BaseSpace)Process(){

}
