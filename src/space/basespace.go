package space

import (
	"fmt"
	"potatoengine/src/entity"
	"potatoengine/src/logService"
	"potatoengine/src/netmessage"
)

type BaseSpace struct {
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
	if _, ok := SpaceMap[id]; ok {
		logService.LogError(fmt.Sprintf("space(id:%s) have same entity(entityid :%s)", this.SpaceID, id))
		return
	}
	this.Entitys[id] = entity
}
