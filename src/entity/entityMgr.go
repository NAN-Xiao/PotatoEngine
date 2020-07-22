package entity

var EntityMap map[int32]IEntity

func init() {
	EntityMap = make(map[int32]IEntity)
}

func GetEntity(conid int32) IEntity {
	entity, ok := EntityMap[conid]
	if ok {
		return entity
	}
	return nil
}
func RegistEntity(entity IEntity) {
	e := entity.GetEntity()
	if &e == nil {
		return
	}
	id := int32(e.EntityID)
	if _, ok := EntityMap[id]; ok == true {
		return
	}
	EntityMap[id] = entity
}
