package entityMgr

import "potatoengine/src/entity"

var EntityMap map[int32]entity.Entity

func GetEntity(conid int32) entity.Entity {
	entity, ok := EntityMap[conid]
	if ok {
		return entity
	}
	return nil
}
func AddEntity(entity entity.Entity)  {
	id:=int32(entity.ConnID)
	if _,ok:=EntityMap[id];ok==true{
		return
	}
	EntityMap[id]=entity
}
