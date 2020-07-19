package entity


var EntityMap map[int32]Entity

func GetEntity(conid int32) Entity {
	entity, ok := EntityMap[conid]
	if ok {
		return entity
	}
	return nil
}
func AddEntity(entity Entity)  {
	id:=int32(entity.Conn.GetID())
	if _,ok:=EntityMap[id];ok==true{
		return
	}
	EntityMap[id]=entity
}
