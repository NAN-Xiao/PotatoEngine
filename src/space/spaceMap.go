package space

var spaceMap map[int32]ISpace

func init() {
	if spaceMap == nil {
		spaceMap = make(map[int32]ISpace, 0)
	}
}
//regist space to globle space map
func AddSpace(sp ISpace) {
	spid := sp.GetSpace().SpaceID
	if _, ok := spaceMap[spid]; ok {
		return
	}
	spaceMap[spid] = sp
}
//根据id查找注册到全局space
func GetSpaceByID(id int32) ISpace {
	if sp, ok := spaceMap[id]; ok {
		return sp
	}
	return nil
}
//根据name查找注册到全局space
func GetSpaceByName(name string) ISpace {
	if spaceMap == nil || len(spaceMap) <= 0 {
		return nil
	}
	for i := range spaceMap {
		if spaceMap[i].GetSpace().Spacename == name {
			return spaceMap[i]
		}
	}
	return nil
}
