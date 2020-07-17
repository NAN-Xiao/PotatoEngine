package entity

type Entity struct {
	id          int32
	spaceid     int32

}

func (this *Entity) Read() interface{} {
}

func (this *Entity) Write(pkg interface{}) {

}

func (this *Entity) GetEntityID() int32 {
	return this.id
}
func (this *Entity) SetEntityID(id int32) {
	this.id = id
}
func (this *Entity) GetSpaceID() int32 {
	return this.spaceid
}

//进入场景
func (this *Entity) EnterSpace(spaceID int32) {

}

//退出场景
func (this *Entity) LeaveSpace(spaceID int32) {

}
