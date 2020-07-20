package space

import (
	"fmt"
	"potatoengine/src/entity"
	"potatoengine/src/logService"
)

type BaseSpace struct {
	GameID     int32
	SpaceID    int32
	Spacename  string
	Entitys    map[int32] entity.IEntity
	//Spacechanl chan netmessage.ServerMsgPackage
}

//从当前space移除entity
func (this *BaseSpace) LeaveSpace(et entity.IEntity) {
	if len(this.Entitys) <= 0 || this.Entitys == nil {
		return
	}
	id := et.GetEntity().EntityID
	_, ok := this.Entitys[id]
	if ok {
		delete(this.Entitys, id)
	}
}

//entity加入当前地图
func (this *BaseSpace) EnterSpace(entity entity.IEntity) {
	if len(this.Entitys) <= 0 || this.Entitys == nil {
		logService.LogError(fmt.Sprintf("space(id:%s) is not have entity map", this.SpaceID))
		return
	}
	id := entity.GetEntity().EntityID
	if sp := GetSpaceByID(id); sp == nil {
		logService.LogError(fmt.Sprintf("space(id:%s) have same entity(entityid :%s)", this.SpaceID, id))
		return
	}
	this.Entitys[id] = entity
}


