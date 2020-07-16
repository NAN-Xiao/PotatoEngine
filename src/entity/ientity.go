package entity

type IEntity interface {
	GetEntityID() int32
	SetEntityID(int32)
	GetSpaceID() int32
	//进入场景
	EnterSpace(spaceID int32)
	//退出场景
	LeaveSpace(spaceID int32)
}
