package space

var SpaceMap map[int32]*BaseSpace

func init() {
	if SpaceMap == nil {
		SpaceMap = make(map[int32]*BaseSpace,0)
	}
}
//regist space to globle space map
func AddSpace(sp *BaseSpace) {
	spid:=sp.SpaceID
	if _,ok:=SpaceMap[spid];ok{
		return
	}
	SpaceMap[spid]=sp
}

func GetSpace(id int32) *BaseSpace {
	if sp,ok:=SpaceMap[id];ok{
		return sp
	}
	return nil
}

