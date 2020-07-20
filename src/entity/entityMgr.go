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
	id := int32(entity.GetEntity().EntityID)
	if _, ok := EntityMap[id]; ok == true {
		return
	}
	EntityMap[id] = entity
}
