package space


var SpaceMap map[int32] ISpace

func init() {
	if SpaceMap == nil {
		SpaceMap = make(map[int32] ISpace,0)
	}
}
//regist space to globle space map
func AddSpace(sp ISpace) {
	spid:=sp.GetID()
	if _,ok:=SpaceMap[spid];ok{
		return
	}
	SpaceMap[spid]=sp
}

func GetSpace(id int32) ISpace {
	if sp,ok:=SpaceMap[id];ok{
		return sp
	}
	return nil
}

